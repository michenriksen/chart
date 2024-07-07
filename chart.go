package chart

import (
	"cmp"
	"errors"
	"fmt"
	"io"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
)

var (
	dataSepRE = regexp.MustCompile(`[\s,;:|#]`) // Matches common separator symbols in tabular data.
	floatRE   = regexp.MustCompile(`[\d\.]`)    // Matches integer and float values.
)

// SortOption represents a sort option for a [Chart].
type SortOption int

// SortDirection represents a sorting direction for a [Chart].
type SortDirection int

const (
	SortNone           SortOption = iota // No sorting.
	SortByLabel                          // Sort by label alphabetically.
	SortByLabelNumeric                   // Sort by label numerically.
	SortByValue                          // Sort by value.
)

const (
	OrderNone SortDirection = iota // No ordering.
	OrderAsc                       // Ascending order.
	OrderDesc                      // Descending order.
)

// Default option values.
const (
	DefaultSort          = SortNone
	DefaultSortDirection = OrderNone
	DefaultPrecision     = 2
)

// Renderer renders a chart to a writer.
type Renderer interface {
	// Render renders the given chart and writes it to the writer.
	Render(*Chart, io.Writer) (int, error)
}

// orderedMap wraps a map of labels and data to record the order of insertion.
type orderedMap struct {
	m  map[string]float64
	k  []string
	mu sync.RWMutex
}

func newOrderedMap() *orderedMap {
	return &orderedMap{m: make(map[string]float64)}
}

func (m *orderedMap) set(key string, val float64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.m[key]; ok {
		m.m[key] = val
		return
	}

	m.k = append(m.k, key)
	m.m[key] = val
}

func (m *orderedMap) get(key string) (float64, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if val, ok := m.m[key]; ok {
		return val, ok
	}

	return 0, false
}

func (m *orderedMap) keys() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	cp := make([]string, len(m.k))
	copy(cp, m.k)

	return cp
}

func (m *orderedMap) values() []float64 {
	m.mu.RLock()
	defer m.mu.RUnlock()

	vals := make([]float64, 0, len(m.k))
	for _, k := range m.k {
		vals = append(vals, m.m[k])
	}

	return vals
}

// Chart represents a simple bar chart.
type Chart struct {
	data    *orderedMap
	sort    SortOption
	sortDir SortDirection
	p       float64
}

// New creates a new [Chart] configured with given options.
func New(opts ...ChartOption) (*Chart, error) {
	c := &Chart{
		data:    newOrderedMap(),
		sort:    DefaultSort,
		sortDir: DefaultSortDirection,
		p:       math.Pow(10, DefaultPrecision),
	}

	for i, opt := range opts {
		if err := opt(c); err != nil {
			return nil, fmt.Errorf("applying option #%d: %w", i+1, err)
		}
	}

	return c, nil
}

// Set sets the value for a label.
func (c *Chart) Set(label string, value float64) *Chart {
	c.data.set(label, value)
	return c
}

// Add adds the number to a label's value.
// If label is not registered, it is added to the chart.
func (c *Chart) Add(label string, value float64) *Chart {
	if val, ok := c.data.get(label); ok {
		value += val
	}

	return c.Set(label, value)
}

// Labels returns chart labels sorted and ordered according to configuration.
func (c *Chart) Labels() []string {
	labels := c.data.keys()

	switch c.sort {
	case SortByLabel:
		slices.SortStableFunc(labels, cmp.Compare)
	case SortByLabelNumeric:
		slices.SortStableFunc(labels, func(i, j string) int {
			return cmp.Compare(stringToInt(i), stringToInt(j))
		})
	case SortByValue:
		slices.SortStableFunc(labels, func(i, j string) int {
			iVal, _ := c.data.get(i)
			jVal, _ := c.data.get(j)

			return cmp.Compare(iVal, jVal)
		})
	}

	if c.sortDir == OrderDesc {
		slices.Reverse(labels)
	}

	return labels
}

// Value returns the value for a label.
// Returns an error if label does not exist.
func (c *Chart) Value(label string) (float64, error) {
	if val, ok := c.data.get(label); ok {
		return math.Round(val*c.p) / c.p, nil
	}

	return 0, errors.New("unknown label")
}

// MaxValue returns the highest chart value.
func (c *Chart) MaxValue() float64 {
	vals := c.data.values()
	if len(vals) == 0 {
		return 0
	}

	maxVal := 0.0
	for _, val := range vals {
		if val > maxVal {
			maxVal = val
		}
	}

	return math.Round(maxVal*c.p) / c.p
}

// MaxLabel returns the longest chart label.
func (c *Chart) MaxLabel() string {
	labels := c.data.keys()
	if len(labels) == 0 {
		return ""
	}

	maxLabel := ""
	for _, label := range labels {
		if len(label) > len(maxLabel) {
			maxLabel = label
		}
	}

	return maxLabel
}

// ChartOption configures a [Chart].
type ChartOption func(*Chart) error

// WithSorting configures a [Chart] with bar sorting options.
func WithSorting(sort SortOption, dir SortDirection) ChartOption {
	return func(c *Chart) error {
		c.sort = sort
		c.sortDir = dir
		return nil
	}
}

// WithPrecision configures a [Chart] with a precision for values.
func WithPrecision(p int) ChartOption {
	return func(c *Chart) error {
		if p < 0 {
			p = 0
		}

		c.p = math.Pow(10, float64(p))
		return nil
	}
}

// ParseLine parses a data line into its float64 value and label string.
//
// The line is expected to have the following structure:
//
//	<numeric value> <label>
//
// The function tolerates any kind of whitespace between the value and label, as
// well as currency symbols and punctuation.
func ParseLine(line string) (float64, string, error) {
	sepIdx := dataSepRE.FindStringIndex(line)
	if sepIdx == nil {
		return 0, "", errors.New("missing data separator")
	}

	value := strings.TrimSpace(line[0:sepIdx[0]])
	label := strings.TrimSpace(line[sepIdx[1]:])

	if label == "" {
		return 0, "", errors.New("missing label")
	}

	value = strings.TrimSuffix(strings.Join(floatRE.FindAllString(value, -1), ""), ".")
	if value == "" {
		return 0, "", errors.New("missing value")
	}

	count, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, "", fmt.Errorf("parsing %q as float64: %w", value, err)
	}

	return count, label, nil
}

// stringToInt strips all non-numeric characters from a string and converts it
// to an integer. Returns 0 if conversion fails.
func stringToInt(s string) int {
	if num, err := strconv.Atoi(s); err == nil {
		return num
	}

	numStr := strings.Join(floatRE.FindAllString(s, -1), "")
	if numStr == "" {
		return 0
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0
	}

	return num
}
