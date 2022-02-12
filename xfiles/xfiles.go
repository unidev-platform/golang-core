package xfiles

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/unidev-platform/golang-core/xcollection"
)

// Distinct - Read text file lines as slice without empty and duplicates
func Distinct(path string) ([]string, error) {
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

	return xcollection.MapKeys(linesMap), scanner.Err()

}

// Find - extract files matching extension
func Find(dir string, extension string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(dir, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == extension {
			files = append(files, s)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
