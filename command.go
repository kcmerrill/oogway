package main

import (
	"bytes"
	"os/exec"
	"strconv"
	"text/template"
	"time"

	sprig "github.com/Masterminds/sprig"
)

type command struct {
	Cmd     string        `yaml:"cmd"`
	After   int           `yaml:"after"`
	Nice    time.Duration `yaml:"nice"`
	Results []byte
	Error   error
	RunTime int64
}

func (c *command) exec(commandType string, check *check) *command {
	// reset the goods
	c.Error, c.Results = nil, nil
	c.RunTime = -1

	// no need to go on really ...
	if !c.okToExec() {
		return c
	}

	// capture when the command started
	started := time.Now().Unix()

	// setup our template
	template, templateError := template.New("cmdTemplate").Funcs(sprig.TxtFuncMap()).Parse(c.Cmd)
	if templateError != nil {
		c.Error = templateError
		return c
	}

	// assign our template to a buffer for later
	b := new(bytes.Buffer)
	templateExecError := template.Execute(b, check)
	if templateExecError != nil {
		c.Error = templateExecError
		return c
	}
	// alright, lets see what we get
	cmd := exec.Command("sh", "-c", b.String())
	c.Results, c.Error = cmd.CombinedOutput()

	// calculate the runtime for the given command
	c.RunTime = time.Now().Unix() - started
	if c.Error == nil {
		return c
	}
	// hmmm ... if we made it here, we error'd out
	return c
}

func (c *command) id() string {
	return c.Cmd + strconv.Itoa(c.After) + c.Nice.String()
}

func (c *command) ok() bool {
	return c.Error == nil
}

func (c *command) okToExec() bool {
	// in case there are additional things we'd want to skip
	// mute, maintenance mode, etc ...
	return !(c.Cmd == "")
}
