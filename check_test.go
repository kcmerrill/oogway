package main

import (
	"os"
	"strings"
	"testing"
	"time"
)

func TestCheck(t *testing.T) {
	oogway := Dojo(
		&oogway{
			ChecksDir:       "t/suiteB",
			CheckInterval:   1 * time.Millisecond,
			ChecksExtension: "oog",
		})

	// load our checks
	oogway.loadChecks()

	// lets pluck out a specific check
	check, exists := oogway.checks["oogway.test"]
	if !exists {
		t.Fatalf("oogway.test is required for this test")
	}

	if check.Attempts != 0 {
		t.Fatalf("Should be starting off on a clean slate. Attempts should be zero")
	}

	// go through 100 checks as fast as you can ...
	for x := 0; x <= 100; x++ {
		check.check()
	}

	// give everything a quick second to catch up(go routines ... )
	<-time.After(time.Second)

	// based on how I have the test setup, the check should be in an OK state now.
	if check.Attempts != 0 {
		t.Fatalf("Should be starting off on a clean slate. Attempts should be zero")
	}

	// Make sure the state is ok
	if strings.ToLower(check.Status) != "ok" {
		t.Fatalf("Expected state to be ok")
	}

	// ok, lets make sure a bunch of files exist ;)
	outputExists := []string{
		"/tmp/oogway.test.check.0",
		"/tmp/oogway.test.check.1",
		"/tmp/oogway.test.check.2",
		"/tmp/oogway.test.check.3",
		"/tmp/oogway.test.check.15",
		"/tmp/oogway.test.warn.1",
		"/tmp/oogway.test.warn.2",
		"/tmp/oogway.test.warn.3",
		"/tmp/oogway.test.warn.4",
		"/tmp/oogway.test.fix.3",
		"/tmp/oogway.test.fix.4",
		"/tmp/oogway.test.fix.5",
		"/tmp/oogway.test.critical.5",
	}
	for _, f := range outputExists {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			// path/to/whatever does not exist
			t.Fatalf("Expected " + string(f) + " to exist")
		}
	}

	// ok, now lets verify that the path is looking ok
	outputShouldNotExist := []string{
		"/tmp/oogway.test.check.30",
		"/tmp/oogway.test.warn.5",
		"/tmp/oogway.test.warn.6",
		"/tmp/oogway.test.fix.2",
		"/tmp/oogway.test.critical.6",
	}
	for _, f := range outputShouldNotExist {
		if _, err := os.Stat(f); !os.IsNotExist(err) {
			// path/to/whatever does not exist
			t.Fatalf("The file " + string(f) + " should not exist")
		}
	}
}
