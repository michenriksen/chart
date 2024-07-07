package cli

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/michenriksen/chart"
	"github.com/michenriksen/chart/simple"
)

const (
	defaultMaxLength      = 80
	defaultMaxLabelLength = 20
	defaultPrecision      = 2
	defaultSort           = "none"
)

//go:embed usage.txt
var usage string

var sortOptMap = map[string]chart.SortOption{
	"none":     chart.SortNone,
	"label":    chart.SortByLabel,
	"labelnum": chart.SortByLabelNumeric,
	"value":    chart.SortByValue,
}

// flags represents the CLI flags.
type flags struct {
	Count          bool   // Count occurrences of lines.
	MaxLength      int    // Maximum chart length.
	MaxLabelLength int    // Maximum label length.
	Precision      int    // Value precision.
	Scale          bool   // Scale bars logarithmically.
	Mermaid        bool   // Create Mermaid XYChart.
	Chartjs        bool   // Create Chart.js configuration.
	Version        bool   // Display version information.
	Title          string // Mermaid chart title.
	in             string
	out            string
	sort           string
	desc           bool
	tick           string
}

// Sort returns the sort option to use.
func (f *flags) Sort() chart.SortOption {
	return sortOptMap[f.sort]
}

// SortDirection returns the sort direction to use.
func (f *flags) SortDirection() chart.SortDirection {
	if f.desc {
		return chart.OrderDesc
	}

	return chart.OrderAsc
}

// Tick returns the tick to use for drawing bars.
func (f *flags) Tick() rune {
	if f.tick == "" {
		return simple.DefaultTick
	}

	return rune(f.tick[0])
}

// In returns the reader to read data from.
// Caller is responsible for closing the reader.
func (f *flags) In() (io.ReadCloser, error) {
	if f.in == "" || f.in == "-" {
		return os.Stdin, nil
	}

	r, err := os.Open(f.in)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}

	return r, nil
}

// Out returns the writer to write chart to.
// Caller is responsible for closing the writer.
func (f *flags) Out() (io.WriteCloser, error) {
	if f.out == "" || f.out == "-" {
		return os.Stdout, nil
	}

	if err := os.MkdirAll(filepath.Dir(f.out), 0o700); err != nil {
		return nil, fmt.Errorf("creating directory: %w", err)
	}

	w, err := os.Create(f.out)
	if err != nil {
		return nil, fmt.Errorf("creating file: %w", err)
	}

	return w, nil
}

// parseFlags parses flags from arguments.
// Returns an error if parsing fails or invalid values are given.
func parseFlags(args []string) (*flags, error) {
	flagset := flag.NewFlagSet("chart", flag.ContinueOnError)
	flagset.Usage = func() {}

	flags := flags{}

	boolFlag(flagset, &flags.Count, "count", "c", false, "count line occurrences")
	intFlag(flagset, &flags.MaxLength, "length", "l", defaultMaxLength, "maximum bar length")
	intFlag(flagset, &flags.MaxLabelLength, "label-length", "L", defaultMaxLabelLength, "maximum label length")
	intFlag(flagset, &flags.Precision, "precision", "p", defaultPrecision, "precision for values")
	boolFlag(flagset, &flags.Scale, "scale", "S", false, "scale bars logarithmically")
	boolFlag(flagset, &flags.Mermaid, "mermaid", "m", false, "create Mermaid XYChart")
	boolFlag(flagset, &flags.Chartjs, "chartjs", "C", false, "create Chart.js configuration")
	boolFlag(flagset, &flags.Version, "version", "v", false, "Display version information and exit")
	stringFlag(flagset, &flags.Title, "title", "T", "", "chart title (mermaid, chartjs)")
	stringFlag(flagset, &flags.in, "in", "i", "", "read data from file")
	stringFlag(flagset, &flags.out, "out", "o", "", "write chart to file")
	stringFlag(flagset, &flags.sort, "sort", "s", defaultSort, "chart sorting option")
	boolFlag(flagset, &flags.desc, "desc", "d", false, "sort chart in descending order")
	stringFlag(flagset, &flags.tick, "tick", "t", "", "use symbol for drawing bars")

	if err := flagset.Parse(args); err != nil {
		return nil, fmt.Errorf("parsing flags: %w", err)
	}

	if _, ok := sortOptMap[flags.sort]; !ok {
		return nil, fmt.Errorf("unknown sort option %q", flags.sort)
	}

	return &flags, nil
}

// printUsage prints application usage to stderr.
// If an error is given, it is printed above the usage.
func printUsage(err error) {
	if err != nil && !errors.Is(err, flag.ErrHelp) {
		fmt.Fprintf(os.Stderr, "Error: %v\n\n", err)
	}

	fmt.Fprintf(os.Stderr, usage, defaultMaxLength, defaultMaxLabelLength, defaultPrecision)
}

func boolFlag(flagset *flag.FlagSet, p *bool, name, short string, value bool, usage string) { //nolint:revive // acceptable arg count.
	flagset.BoolVar(p, name, value, usage)
	if short != "" {
		flagset.BoolVar(p, short, value, usage)
	}
}

func intFlag(flagset *flag.FlagSet, p *int, name, short string, value int, usage string) { //nolint:revive // acceptable arg count.
	flagset.IntVar(p, name, value, usage)
	if short != "" {
		flagset.IntVar(p, short, value, usage)
	}
}

func stringFlag(flagset *flag.FlagSet, p *string, name, short, value, usage string) { //nolint:revive // acceptable arg count.
	flagset.StringVar(p, name, value, usage)
	if short != "" {
		flagset.StringVar(p, short, value, usage)
	}
}
