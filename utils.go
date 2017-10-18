package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func combineConfigFiles(dir, extension string) []byte {
	files, filesError := filepath.Glob(strings.TrimRight(dir, "/") + "/*." + extension)
	if filesError != nil {
		return []byte{}
	}

	config := []byte{}
	for _, file := range files {
		contents, _ := ioutil.ReadFile(file)
		config = append(config, []byte("\n")...)
		config = append(config, contents...)
	}

	return config
}
