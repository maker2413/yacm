package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	versionCmd := NewVersionCmd()
	assert.NotNil(t, versionCmd)

	t.Run("printVersion - No parent", func(t *testing.T) {
		err := printVersion(versionCmd, []string{})
		assert.Error(t, err)
		assert.Equal(t, err.Error(), "Unable to read Parent()")
	})

	child := NewVersionCmd()
	versionCmd.AddCommand(child)

	t.Run("printVersion - No version", func(t *testing.T) {
		err := printVersion(child, []string{})
		assert.Error(t, err)
		assert.Equal(t, err.Error(), "Unable to read Parent().Version")
	})

	versionCmd.Version = "v1"

	t.Run("printVersion", func(t *testing.T) {
		err := printVersion(child, []string{})
		assert.NoError(t, err)
	})
}
