package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoot(t *testing.T) {
	rootCmd := NewRootCmd("test")
	assert.NotNil(t, rootCmd)

	assert.NotPanics(t, initConfig)
}
