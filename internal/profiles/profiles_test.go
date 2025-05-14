package profiles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfiles(t *testing.T) {
	p := Profiles{}
	assert.NotNil(t, p)

	t.Run("loadProfiles", func(t *testing.T) {
		err := p.loadProfiles()
		assert.NoError(t, err)
	})
}
