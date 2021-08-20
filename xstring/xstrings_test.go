package xstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringExtractions(t *testing.T) {
	items := ExtractStringBetween(" 1 qwe 2  1 xxx 2", "1", "2")

	assert.Equal(t, 2, len(items))

}