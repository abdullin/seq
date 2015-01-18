package seq

import "fmt"

type Map map[string]interface{}

type Result struct {
	Diffs []string
}

func (r *Result) Ok() bool {
	return len(r.Diffs) == 0
}

func (m Map) Test(subj interface{}) *Result {
	expected := flatten("", objectToMap(m))
	actual := flatten("", objectToMap(subj))
	//fmt.Println("expected")
	//debug(expected)
	//fmt.Println("actual")
	//debug(actual)
	diffs := diff(expected, actual)
	return &Result{diffs}
}

func debug(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, ":", v)
	}
}
