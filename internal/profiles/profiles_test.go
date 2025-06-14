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

	p, err := Init()
	assert.NoError(t, err)
	assert.NotNil(t, p)

	t.Run("GetProfiles", func(t *testing.T) {
		expected := make(map[string]SystemProfile)
		expected["test.yml"] = SystemProfile{}

		assert.Equal(t, expected, p.GetProfiles())
	})

	t.Run("Exist", func(t *testing.T) {
		assert.True(t, p.Exist())
	})

	t.Run("Exist - false", func(t *testing.T) {
		viper.Set("yacm_profiles_dir", "./tests/noprofiles/")

		p2, err := Init()
		assert.NoError(t, err)
		assert.False(t, p2.Exist())
	})

	t.Run("Exist - false from error", func(t *testing.T) {
		viper.Set("yacm_profiles_dir", "./tests/nonexistent/")

		p2, err := Init()
		assert.Error(t, err)
		assert.False(t, p2.Exist())
	})
}
