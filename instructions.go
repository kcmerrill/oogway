package main

import (
	b64 "encoding/base64"
	"strconv"
	"time"
)

type instructions struct {
	Summary  string        `yaml:"summary"`
	Every    time.Duration `yaml:"every"`
	Tries    int           `yaml:"try"`
	Reset    time.Duration `yaml:"reset"`
	Check    command       `yaml:"check"`
	OK       command       `yaml:"ok"`
	Warning  command       `yaml:"warning"`
	Critical command       `yaml:"critical"`
	Fix      command       `yaml:"fix"`
	Recover  command       `yaml:"recover"`
}

func (i *instructions) Try() int {
	if i.Tries == 0 {
		return 1
	}
	return i.Tries
}
func (i *instructions) id() string {
	return b64.StdEncoding.EncodeToString([]byte(i.Summary + strconv.Itoa(i.Tries) + i.Check.id() + i.OK.id() + i.Warning.id() + i.Critical.id() + i.Fix.id() + i.Recover.id()))
}

func (i *instructions) every() time.Duration {
	if i.Every == (0 * time.Second) {
		return (30 * time.Second)
	}
	return i.Every
}
