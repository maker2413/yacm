package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBootstrap(t *testing.T) {
	bootstrapCmd := NewBootstrapCmd()
	assert.NotNil(t, bootstrapCmd)

	t.Run("boostrap", func(t *testing.T) {
		err := bootstrap(bootstrapCmd, []string{})
		assert.NoError(t, err)
	})
}
