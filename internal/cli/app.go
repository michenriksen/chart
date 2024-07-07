package cli

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/michenriksen/chart"
	"github.com/michenriksen/chart/chartjs"
	"github.com/michenriksen/chart/mermaid"
	"github.com/michenriksen/chart/simple"
)

const (
	exitNormal = iota
	exitError
)

const versionTmpl = `chart:
  Version: %s
  Commit:  %s
  Time:    %s
`

func Run() int {
	initLogger()

	flags, err := parseFlags(os.Args[1:])
	if err != nil {
		printUsage(err)

		if errors.Is(err, flag.ErrHelp) {
			return exitNormal
		}

		return exitError
	}

	if flags.Version {
		fmt.Fprintf(os.Stdout, versionTmpl, version, BuildRevision(), BuildTime().Format(time.RFC3339))
		return exitNormal
	}

	c, err := chart.New(
		chart.WithSorting(flags.Sort(), flags.SortDirection()),
		chart.WithPrecision(flags.Precision),
	)
	if err != nil {
		return fatal("creating chart", err)
	}

	in, err := flags.In()
	if err != nil {
		return fatal("opening input", err)
	}

	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if flags.Count {
			c.Add(strings.TrimSpace(scanner.Text()), 1)
			continue
		}

		value, label, err := chart.ParseLine(line)
		if err != nil {
			slog.Warn("skipping unparsable line", "error", err, "line", line)
			continue
		}

		c.Set(label, value)
	}

	in.Close()

	var renderer chart.Renderer

	switch {
	case flags.Mermaid:
		renderer, err = mermaid.NewRenderer(
			mermaid.WithTitle(flags.Title),
		)
	case flags.Chartjs:
		renderer, err = chartjs.NewRenderer(
			chartjs.WithTitle(flags.Title),
		)
	default:
		renderer, err = simple.NewRenderer(
			simple.WithMaxLength(flags.MaxLength),
			simple.WithMaxLabelLength(flags.MaxLabelLength),
			simple.WithScaling(flags.Scale),
			simple.WithTick(flags.Tick()),
		)
	}

	if err != nil {
		return fatal("creating renderer", err)
	}

	out, err := flags.Out()
	if err != nil {
		return fatal("opening output", err)
	}
	defer out.Close()

	if _, err := renderer.Render(c, out); err != nil {
		return fatal("rendering chart", err)
	}

	return exitNormal
}

func initLogger() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	})))
}

func fatal(message string, err error, args ...any) int {
	if err != nil {
		args = append(args, "error", err)
	}

	slog.Error(message, args...)
	return exitError
}
