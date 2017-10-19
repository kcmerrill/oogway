package main

import (
	"os/exec"
	"strconv"
	"time"
)

type command struct {
	Cmd     string        `yaml:"cmd"`
	After   int           `yaml:"after"`
	Every   time.Duration `yaml:"every"`
	results []byte
	error   error
	RunTime int64
	Reset   time.Duration `yaml:"reset"`
}

func (c *command) exec() *command {
	// reset the goods
	c.error, c.results = nil, nil

	// no need to go on really ...
	if c.Cmd == "" {
		return c
	}

	// capture when the command started
	started := time.Now().Unix()

	// alright, lets see what we get
	cmd := exec.Command("sh", "-c", c.Cmd)
	c.results, c.error = cmd.CombinedOutput()

	// calculate the runtime for the given command
	c.RunTime = time.Now().Unix() - started
	return c
}

func (c *command) id() string {
	return c.Cmd + c.Summary + strconv.Itoa(c.After) + strconv.Itoa(c.Tries) + c.Every.String() + c.Reset.String()
}

func (c *command) ok() bool {
	return c.error == nil
}
