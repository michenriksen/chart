package main_test

import (
	"os"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"

	"github.com/michenriksen/chart/internal/cli"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"chart": cli.Run,
	}))
}

func Test(t *testing.T) {
	_, updateGolden := os.LookupEnv("TEST_UPDATE_GOLDEN")

	testscript.Run(t, testscript.Params{
		Dir:           "testdata/script",
		UpdateScripts: updateGolden,
	})
}
