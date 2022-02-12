package xenv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchingCurrentUser(t *testing.T) {
	t.Parallel()
	user := Env("USER", "")

	assert.NotEmpty(t, user)
}
