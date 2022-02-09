package xenv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchingCurrentUser(t *testing.T) {

	user := Env("USER", "")

	assert.NotEmpty(t, user)
}
