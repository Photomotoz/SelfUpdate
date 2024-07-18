package main

import (
	"path/filepath"
)

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func parsePath(path string) (string, string, string) {
	absolutePath, err := filepath.EvalSymlinks(path)
	errorCheck(err)
	return absolutePath, filepath.Dir(absolutePath), filepath.Base(absolutePath)
}
