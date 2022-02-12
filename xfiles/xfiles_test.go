package xfiles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileLinesExtraction(t *testing.T) {
	t.Parallel()
	lines, err := Distinct("distinct_file_lines_test.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 6, len(lines))
}

func TestListFiles(t *testing.T) {
	t.Parallel()
	files, err := Find(".", ".txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 1, len(files))
}
