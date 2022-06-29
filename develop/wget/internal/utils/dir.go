package utils

import (
	"log"
	"os"
	"strings"
)

// make directory for site

func MakeDir(path string, new string) {
	dir := MakePath(path, new)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func MakePath(args ...string) string {
	return strings.Join(args, "/")
}
