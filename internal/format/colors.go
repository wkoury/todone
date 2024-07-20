package format

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

func colorizePatternInLine(line string, substr string) string {
	re, err := regexp.Compile(substr)
	if err != nil {
		panic(err)
	}

	return re.ReplaceAllStringFunc(line, func(match string) string {
		return color.RedString(match)
	})
}

func WriteMsg(path string, lineIndex int, line string, substr string) {
	fmt.Printf("%s line %d:\n", path, lineIndex+1)
	fmt.Printf("\t%s\n", colorizePatternInLine(strings.TrimSpace(line), substr))
}
