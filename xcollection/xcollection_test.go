package xcollection

import (
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestStringBoolMapKeys(t *testing.T) {

	keys := StringBoolMapKeys(map[string]bool{
		"qwe": true,
		"123": true,
	})

	assert.Equal(t, 2, len(keys))
}

func TestRandomSelection(t *testing.T) {
	rand.Seed(time.Now().Unix())
	element := StringRandomElement([]string{"1", "2", "3", "4", "5", "6"})
	log.Printf("Random element: %s", element)
	assert.NotNil(t, element)
}
