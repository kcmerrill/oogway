package main

import (
	"strings"
	"testing"
)

func TestUtilsCombineConfigFiles(t *testing.T) {
	// just so happens the files are .com's :D
	files := combineConfigFiles("t/suiteA/", "com")
	// you should find BOTH kcmerrill.com and google.com
	if !strings.Contains(string(files), "kcmerrill.com") {
		t.Fatalf("kcmerrill.com was expected. Not found.")
	}

	// google is in a diffferent file
	if !strings.Contains(string(files), "google.com") {
		t.Fatalf("google.com was expected. Not found.")
	}
}
