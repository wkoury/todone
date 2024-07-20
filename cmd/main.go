package main

import (
	"os"

	"github.com/wkoury/todone/internal/files"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	err = files.SearchDirectory(currentDir)
	if err != nil {
		panic(err)
	}
}
