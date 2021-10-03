package xenv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchingCurrentUser(t *testing.T) {

	user := Env("USER", "")

	assert.NotEmpty(t, user)
}
