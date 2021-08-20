package xcollection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringBoolMapKeys(t *testing.T) {

	keys := StringBoolMapKeys(map[string]bool{
		"qwe": true,
		"123": true,
	})

	assert.Equal(t, 2, len(keys))
}
