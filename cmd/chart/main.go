package main

import (
	"os"

	"github.com/michenriksen/chart/internal/cli"
)

func main() {
	os.Exit(cli.Run())
}
