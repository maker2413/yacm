package profiles

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestProfiles(t *testing.T) {
	t.Run("Init - error", func(t *testing.T) {
		_, err := Init()
		assert.Error(t, err)
	})

	viper.Set("yacm_profiles_dir", "./tests/profiles/")

	sp, err := Init()
	assert.NoError(t, err)
	assert.NotNil(t, sp)

	t.Run("exists", func(t *testing.T) {
		assert.True(t, sp.exists())
	})

	t.Run("exists - false", func(t *testing.T) {
		viper.Set("yacm_profiles_dir", "./tests/noprofiles/")

		sp2, err := Init()
		assert.NoError(t, err)
		assert.False(t, sp2.exists())
	})

	t.Run("exists - false from error", func(t *testing.T) {
		viper.Set("yacm_profiles_dir", "./tests/nonexistent/")

		sp2, err := Init()
		assert.Error(t, err)
		assert.False(t, sp2.exists())
	})
}
