package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	initCmd := NewInitCmd()
	assert.NotNil(t, initCmd)

	t.Run("executeInit", func(t *testing.T) {
	})
}
