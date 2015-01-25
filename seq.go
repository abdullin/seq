package seq

import "fmt"

func Test(expected, actual interface{}) *Result {
	return TestEx(expected, actual, nil)
}

type Groups map[interface{}]string

func TestEx(expected, actual interface{}, groups Groups) *Result {

	eMap := flatten("", objectToMap(expected))
	aMap := flatten("", objectToMap(actual))
	result := diff(eMap, aMap, groups)
	return result
}

type Map map[string]interface{}

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
