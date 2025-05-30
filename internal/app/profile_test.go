package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfile(t *testing.T) {
	profileCmd := NewProfileCmd()
	assert.NotNil(t, profileCmd)

	t.Run("executeProfile", func(t *testing.T) {
		err := executeProfile(profileCmd, []string{})
		assert.NoError(t, err)
	})
}
