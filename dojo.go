package main

import (
	"strings"
	"sync"
)

// Dojo creates a place for oogway to get setup
func Dojo(o *oogway) *oogway {
	// lets init some basics here
	o.checks = make(map[string]*check)
	o.checksLock = &sync.Mutex{}
	o.ChecksDir = strings.TrimRight(o.ChecksDir, "/") + "/"

	return o
}
