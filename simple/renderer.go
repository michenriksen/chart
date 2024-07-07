package simple

import (
	"errors"
	"fmt"
	"io"
	"math"
	"strings"
	"unicode/utf8"

	"github.com/michenriksen/chart"
)

const smallTick = '▏'

// Default option values.
const (
	DefaultTick           = '▇'
	DefaultMaxLength      = 80
	DefaultMaxLabelLength = 20
	DefaultScale          = false
)

// Renderer renders a [chart.Chart] with simple characters and symbols suitable
// for display in terminals and text files.
type Renderer struct {
	maxLen          int
	maxLabelLen     int
	scale           bool
	tick            rune
	longestLabelLen int
	longestValLen   int
	maxVal          float64
	barLen          int
}

// NewRenderer returns a [chart.Renderer] for rendering a [chart.Chart] with
// simple characters and symbols suitable for display in terminals and text
// files.
func NewRenderer(opts ...RendererOption) (*Renderer, error) {
	r := &Renderer{
		maxLen:      DefaultMaxLength,
		maxLabelLen: DefaultMaxLabelLength,
		scale:       DefaultScale,
		tick:        DefaultTick,
	}

	for i, opt := range opts {
		if err := opt(r); err != nil {
			return nil, fmt.Errorf("applying option #%d: %w", i+1, err)
		}
	}

	return r, nil
}

// Render renders chart to out writer.
func (r *Renderer) Render(c *chart.Chart, out io.Writer) (int, error) {
	r.maxVal = c.MaxValue()
	r.longestLabelLen = min(len(c.MaxLabel()), r.maxLabelLen)
	r.longestValLen = len(r.value(r.maxVal))
	r.barLen = r.maxLen - r.longestLabelLen - r.longestValLen - 2

	written := 0

	for _, label := range c.Labels() {
		value, err := c.Value(label)
		if err != nil {
			return written, fmt.Errorf("getting value for %q label: %w", label, err)
		}

		n, err := r.write(label, value, out)
		if err != nil {
			return written, fmt.Errorf("writing bar for label %q (value %g): %w", label, value, err)
		}

		written += n
	}

	return written, nil
}

func (r *Renderer) write(label string, value float64, out io.Writer) (int, error) {
	n, err := fmt.Fprintf(out, "%s %s %s\n", r.label(label), r.bar(value), r.value(value))
	if err != nil {
		return n, fmt.Errorf("writing to out: %w", err)
	}

	return n, nil
}

func (r *Renderer) bar(value float64) string {
	length := value / r.maxVal * float64(r.barLen)
	if r.scale {
		length = math.Log10(value+1) / math.Log10(float64(r.maxVal)+1) * float64(r.barLen)
	}
	length = math.Round(length)

	if length == 0 {
		if r.tick == DefaultTick {
			return string(smallTick)
		}

		return ""
	}

	return strings.Repeat(string(r.tick), int(length))
}

func (r *Renderer) label(label string) string {
	if len(label) > r.maxLabelLen {
		label = truncate(label, r.maxLabelLen)
	}

	format := fmt.Sprintf("%%%ds", min(r.longestLabelLen, r.maxLabelLen))

	return fmt.Sprintf(format, label)
}

func (*Renderer) value(value float64) string {
	return fmt.Sprintf("%g", value)
}

// RendererOption configures a [Renderer].
type RendererOption func(*Renderer) error

// WithMaxLength configures a [Renderer] with a maximum chart length.
func WithMaxLength(n int) RendererOption {
	return func(r *Renderer) error {
		if n <= 0 {
			return errors.New("maximum length must be a positive integer")
		}

		r.maxLen = n
		return nil
	}
}

// WithMaxLabelLength configures a [Renderer] with a maximum label length.
// If a label exceeds the maximum length, it will be truncated in the middle.
func WithMaxLabelLength(n int) RendererOption {
	return func(r *Renderer) error {
		if n <= 0 {
			return errors.New("maximum label length must be a positive integer")
		}

		r.maxLabelLen = n
		return nil
	}
}

// WithScaling configures a [Renderer] to scale chart bars logarithmically.
func WithScaling(enable bool) RendererOption {
	return func(r *Renderer) error {
		r.scale = enable
		return nil
	}
}

// WithTick configures a [Renderer] with a rune to use for drawing chart bars.
func WithTick(tick rune) RendererOption {
	return func(r *Renderer) error {
		r.tick = tick
		return nil
	}
}

func truncate(s string, maxLen int) string {
	sLen := utf8.RuneCountInString(s)
	if sLen <= maxLen {
		return s
	}

	ellips := "..."
	ellipsLen := len(ellips)

	partLen := (maxLen - ellipsLen) / 2
	start := s[:partLen]
	end := s[sLen-partLen:]

	return start + ellips + end
}
