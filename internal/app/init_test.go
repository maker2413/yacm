package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	initCmd := NewInitCmd()
	assert.NotNil(t, initCmd)

	t.Run("executeInit - error", func(t *testing.T) {
		err := initCmd.Execute()
		assert.Error(t, err)
	})
}
