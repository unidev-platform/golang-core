package xfiles

import (
	"bufio"
	"github.com/unidev-platform/golang-core/xcollection"
	"os"
	"strings"
)

// ReadDistinctFileLines - Read text file lines as slice without empty and duplicates
func ReadDistinctFileLines(path string)([]string, error) {
	var linesMap = make(map[string]bool)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			linesMap[line] = true
		}
	}

	return xcollection.StringBoolMapKeys(linesMap), scanner.Err()

}
