package main

import (
	"os"

	"github.com/maker2413/yacm/cmd"
)

const version string = "v0.3.0"

func main() {
	err := cmd.Execute(version)
	if err != nil {
		os.Exit(1)
	}
}
