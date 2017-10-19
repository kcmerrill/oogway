package main

import (
	"testing"
	"time"
)

func TestDojo(t *testing.T) {
	oogway := Dojo(
		&oogway{
			ChecksDir:       "t/suiteA",
			CheckInterval:   0 * time.Second,
			ChecksExtension: "com",
		})

	// dojo doesn't do a whole lot. yet.
	// biggest thing to check is the checksDir
	// notice the extra slash
	if oogway.ChecksDir != "t/suiteA/" {
		t.Fatalf("Expected a trailing slash at the end of checks directory")
	}
}
