package app

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfigCrashes(t *testing.T) {
	if os.Getenv("TEST") == "1" {
		t.Run("nonexistent config file", func(t *testing.T) {
			cfgFile = "./doesntexist.yml"

			initConfig()
			return
		})
	}

	if os.Getenv("TEST") == "2" {
		t.Run("write config file error", func(t *testing.T) {
			t.Setenv("YACM_BASE_DIR", "./nonexistentdir")

			initConfig()
			return
		})
	}

	expectedErrorString := "exit status 1"

	cmd := exec.Command(os.Args[0], "-test.run=TestInitConfigCrashes")
	cmd.Env = append(os.Environ(), "TEST=1")

	err := cmd.Run()

	// Cast the error as *exec.ExitError and compare the result
	e, ok := err.(*exec.ExitError)
	assert.Equal(t, true, ok)
	assert.Equal(t, expectedErrorString, e.Error())

	cmd = exec.Command(os.Args[0], "-test.run=TestInitConfigCrashes")
	cmd.Env = append(os.Environ(), "TEST=2")

	err = cmd.Run()

	// Cast the error as *exec.ExitError and compare the result
	e, ok = err.(*exec.ExitError)
	assert.Equal(t, true, ok)
	assert.Equal(t, expectedErrorString, e.Error())
}
