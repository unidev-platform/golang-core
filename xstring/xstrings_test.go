package xstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringExtractions(t *testing.T) {
	items := Between(" 1 qwe 2  1 xxx 2", "1", "2")

	assert.Equal(t, 2, len(items))

}
