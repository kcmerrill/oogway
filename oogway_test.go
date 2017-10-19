package main

import (
	"log"
	"testing"
	"time"
)

func TestOogwayLoadChannel(t *testing.T) {
	oogway := Dojo(
		&oogway{
			ChecksDir:       "t/suiteA",
			CheckInterval:   0 * time.Second,
			ChecksExtension: "com",
		})

	if oogway.loadChecks() != nil {
		log.Fatalf("There was a problem unmarshling the YAML")
	}

	// kcmerrill.com check should exist
	if check, exists := oogway.checks["kcmerrill.com"]; exists {
		// checks should be initalized
		if check.Attempts != 0 {
			log.Fatalf("kcmerrill.com has not been attempted. Should be zero")
		}

		// last checked
		if !check.LastChecked.IsZero() {
			log.Fatalf("kcmerill.com last checked should be never")
		}

		// the check command
		if check.Instructions.Check.Cmd != "touch /tmp/suiteA.check.kcmerrill.com" {
			log.Fatalf("kcmerrill.com command should have been set")
		}

		// when to reset
		if check.Instructions.Check.Reset != time.Hour {
			log.Fatalf("Reset should have been set to 1h")
		}

		// when to go critical
		if check.Instructions.Check.Tries != 4 {
			log.Fatalf("Try should have been set to 4 before going critical")
		}
	} else {
		log.Fatalf("kcmerrill.com check should exist. Not found.")
	}
}
