package app

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	if os.Getenv("TEST") == "1" {
		cfgFile = "./doesntexist.yml"
		initConfig()
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestInitConfig")
	cmd.Env = append(os.Environ(), "TEST=1")

	err := cmd.Run()

	// Cast the error as *exec.ExitError and compare the result
	e, ok := err.(*exec.ExitError)
	expectedErrorString := "exit status 1"
	assert.Equal(t, true, ok)
	assert.Equal(t, expectedErrorString, e.Error())
}
