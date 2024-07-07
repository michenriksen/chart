package mermaid

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/michenriksen/chart"
)

// Renderer renders a [chart.Chart] as a Mermaid XYChart.
//
// See: https://mermaid.js.org/syntax/xyChart.html
type Renderer struct {
	title string
}

// NewRenderer returns a [chart.Renderer] for rendering a [chart.Chart] as a
// Mermaid XYChart.
//
// See: https://mermaid.js.org/syntax/xyChart.html
func NewRenderer(opts ...RendererOption) (*Renderer, error) {
	r := &Renderer{}

	for i, opt := range opts {
		if err := opt(r); err != nil {
			return nil, fmt.Errorf("applying option #%d: %w", i+1, err)
		}
	}

	return r, nil
}

// Render renders chart to out writer.
func (r *Renderer) Render(c *chart.Chart, out io.Writer) (int, error) {
	labels := c.Labels()
	values := make([]string, 0, len(labels))

	for _, label := range labels {
		value, err := c.Value(label)
		if err != nil {
			return 0, fmt.Errorf("getting value for %q label: %w", label, err)
		}

		values = append(values, fmt.Sprintf("%g", value))
	}

	buf := new(bytes.Buffer)

	fmt.Fprintln(buf, "xychart-beta")

	if r.title != "" {
		fmt.Fprintf(buf, "  title \"%s\"\n", r.title)
	}

	fmt.Fprintf(buf, "  x-axis [\"%s\"]\n", strings.Join(labels, `", "`))
	fmt.Fprintf(buf, "  bar [%s]\n", strings.Join(values, ", "))

	n, err := out.Write(buf.Bytes())
	if err != nil {
		return n, fmt.Errorf("writing to out: %w", err)
	}

	return n, nil
}

// RendererOption configures a [Renderer].
type RendererOption func(*Renderer) error

// WithTitle configures a [Renderer] with a chart title.
func WithTitle(title string) RendererOption {
	return func(r *Renderer) error {
		r.title = title
		return nil
	}
}
