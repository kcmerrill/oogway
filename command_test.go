package main

import "testing"

func TestCommand(t *testing.T) {
	c := &command{
		Cmd: "echo 'Hellow' > /tmp/TestCommand.txt",
	}

	// check exec status
	c.exec()

	if c.error != nil {
		t.Fatalf("Not expecting c.exec() to throw an error")
	}
}
