package xcollection

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStringBoolMapKeys(t *testing.T) {
	t.Parallel()
	keys := MapKeys(map[string]bool{
		"qwe": true,
		"123": true,
	})

	assert.Equal(t, 2, len(keys))
}

func TestRandomSelection(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().Unix())
	element := Random([]string{"1", "2", "3", "4", "5", "6"})
	log.Printf("Random element: %s", element)
	assert.NotNil(t, element)
}
