package main

import (
	"os"

	"github.com/maker2413/yacm/internal/app"
)

const version string = "v0.3.0"

func main() {
	err := app.Execute(version)
	if err != nil {
		os.Exit(1)
	}
}
