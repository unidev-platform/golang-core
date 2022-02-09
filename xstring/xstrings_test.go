package xstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringExtractions(t *testing.T) {
	items := Between(" 1qwe2 666 1xxx2 000", "1", "2")

	assert.Equal(t, 2, len(items))
	assert.Contains(t, items, "qwe")
	assert.Contains(t, items, "xxx")

}
