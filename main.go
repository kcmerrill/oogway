package main

import (
	"flag"
	"time"
)

func main() {
	checks := flag.String("check-dir", ".", "Directory of checks")
	checksExtension := flag.String("check-extension", "oog", "Extension of check yaml files")
	checkInterval := flag.Duration("check-interval", time.Second, "How often the checks should be reloaded?")
	flag.Parse()

	// Welcome to the Oogway's dojo
	Dojo(&oogway{
		ChecksDir:       *checks,
		CheckInterval:   *checkInterval,
		ChecksExtension: *checksExtension,
	}).instruct()
}
