package seq

import "fmt"

type Result struct {
	Diffs []Diff
}

type Diff struct {
	Path          string
	ExpectedValue string
	ActualValue   string
}

func (r *Result) Ok() bool {
	return len(r.Diffs) == 0
}

func (d *Diff) String() string {
	return fmt.Sprintf("Expected %s to be '%v' but got %s",
		d.Path,
		d.ExpectedValue,
		d.ExpectedValue,
	)
}

func NewResult() *Result {
	return &Result{}
}

func (r *Result) AddDiff(key, expected, actual string) {
	r.Diffs = append(r.Diffs, Diff{key, expected, actual})
}
