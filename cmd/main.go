package main

import (
	"os"
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
