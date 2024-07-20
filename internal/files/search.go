package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/wkoury/todone/internal/format"
)

func GetLines(path string) []string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("error", err)
	}
	contents := string(bytes)

	return strings.Split(contents, "\n")
}

func SearchFile(path string) {
	lines := GetLines(path)

	for ii, line := range lines {
		// If the line is more than 200 characters long, it probably isn't something worth pointing out
		if len(line) <= 200 {
			if strings.Contains(line, "TODO") {
				format.WriteMsg(path, ii+1, line, "TODO")
			}
			if strings.Contains(line, "FIXME") {
				format.WriteMsg(path, ii+1, line, "FIXME")
			}
		}
	}
}

func SearchDirectory(path string) error {
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// TODO: find more ignore paths here
		if !info.IsDir() && !strings.Contains(path, "node_modules") {
			SearchFile(path)
		}

		return nil
	})
}
