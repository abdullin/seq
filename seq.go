package seq

import "fmt"

func Test(expected, actual interface{}) *Result {
	eMap := flatten("", objectToMap(expected))
	aMap := flatten("", objectToMap(actual))
	//fmt.Println("expected")
	//debug(expected)
	//fmt.Println("actual")
	//debug(actual)
	diffs := diff(eMap, aMap)
	return &Result{diffs}
}

type Map map[string]interface{}

type Result struct {
	Diffs []string
}

func (r *Result) Ok() bool {
	return len(r.Diffs) == 0
}

func (m Map) Test(actual interface{}) *Result {
	return Test(m, actual)
}

func debug(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, ":", v)
	}
}
