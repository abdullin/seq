package seq

import "fmt"

type Result struct {
	Diffs []string

	Captures map[string][]Capture
}

func NewResult() *Result {
	return &Result{
		Captures: map[string][]Capture{},
	}
}

type Capture struct {
	Path   string
	Actual string
}

func (r *Result) Capture(group, key, actual string) {
	r.Captures[group] = append(r.Captures[group], Capture{key, actual})
}

func (r *Result) AddDiff(key, expected, actual string) {

	var res = fmt.Sprintf("Expected %s to be '%v' but got %s", key, expected, actual)
	r.Diffs = append(r.Diffs, res)

}
