package profiles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfiles(t *testing.T) {
	p := Profiles{}
	assert.NotNil(t, p)

	t.Run("GetProfiles", func(t *testing.T) {
		err := p.GetProfiles()
		assert.NoError(t, err)
	})
}
