package app

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	initCmd := NewInitCmd()
	assert.NotNil(t, initCmd)

	t.Run("executeInit - error", func(t *testing.T) {
		err := initCmd.Execute()
		assert.Error(t, err)
	})

	t.Run("executeInit", func(t *testing.T) {
		viper.Set("yacm_profiles_dir", "./tests/profiles/")

		err := initCmd.Execute()
		assert.NoError(t, err)
	})
}
