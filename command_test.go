package main

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	c := &command{
		Cmd: "echo 'Hellow' > /tmp/TestCommand.txt",
	}

	// check exec status
	c.exec("name", &check{})

	if c.Error != nil {
		t.Fatalf("Not expecting c.exec() to throw an error")
	}
}

func TestCommandID(t *testing.T) {
	c := &command{
		Cmd:   "1234",
		After: 1234,
	}

	if c.id() != "12341234" {
		fmt.Println(c.id())
		t.Fatalf("Command id() did not return the correct id()")
	}
}

func TestCommandOKToExec(t *testing.T) {
	c := &command{
		Cmd:   "1234",
		After: 1234,
	}

	if c.okToExec() == false {
		t.Fatalf("There is a command, and muted is not set, this should be ok to exec.")
	}
}
