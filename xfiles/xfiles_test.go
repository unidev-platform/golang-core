package xfiles

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileLinesExtraction(t *testing.T) {
	lines, err := ReadDistinctFileLines("distinct_file_lines_test.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 6, len(lines))
}
