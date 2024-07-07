package cli

import (
	"fmt"
	"runtime/debug"
	"strings"
	"time"
)

var buildInfo *debug.BuildInfo

var version = "0.1.0"

// Version returns the current version of the CLI application.
func Version() string {
	if getBIKey("vcs.modified", "false") == "true" {
		return version + "-next"
	}

	return version
}

// BuildRevision returns the current git commit hash.
func BuildRevision() string {
	return getBIKey("vcs.revision", strings.Repeat("0", 40))
}

// BuildShortRevision returns the current shorthand git commit hash.
func BuildShortRevision() string {
	return BuildRevision()[:8]
}

// BuildTime returns the commit time in UTC.
func BuildTime() time.Time {
	ts := getBIKey("vcs.time", time.Now().Format(time.RFC3339))
	t, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		panic(fmt.Errorf("parsing vcs.time in BuildInfo: %w", err))
	}

	return t.UTC()
}

func getBIKey(name, defaultVal string) string {
	if buildInfo == nil {
		var ok bool

		buildInfo, ok = debug.ReadBuildInfo()
		if !ok {
			return defaultVal
		}
	}

	for _, bs := range buildInfo.Settings {
		if bs.Key == name {
			if bs.Value == "" {
				return defaultVal
			}

			return bs.Value
		}
	}

	return defaultVal
}
