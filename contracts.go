package seq

import (
	"fmt"
	"strconv"
	"strings"
)

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
func hasNestedObject(actual map[string]string, key string) bool {
	for k, _ := range actual {
		if strings.HasPrefix(k, key) {
			return true
		}
	}
	return false
}

func diff(expected, actual map[string]string) []string {
	var res []string

	for ek, ev := range expected {
		var av, ok = actual[ek]
		var diff = ""

		if !ok {
			if hasNestedObject(actual, ek) {
				diff = "{Object}"
			} else {
				diff = "nothing"
			}

		} else if av != ev {
			switch ev {
			case "seq:ignore":
				break
			default:
				diff = fmt.Sprintf("'%s'", av)
			}
		}

		if diff != "" {
			res = append(res, fmt.Sprintf("Expected %s to be '%v' but got %s", ek, ev, diff))
		}
	}

	return res
}

func debug(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, ":", v)
	}
}

func propertyPath(key string, name string) string {
	if key == "" {
		return name
	}
	return key + "." + name
}

func flatten(key string, x interface{}) map[string]string {

	var res = make(map[string]string)
	switch vv := x.(type) {
	case string:
		res[key] = vv
	case []interface{}:
		res[fmt.Sprintf("%s.len", key)] = strconv.Itoa(len(vv))
		for ii, iv := range vv {
			var prefix = fmt.Sprintf("%s[%v]", key, ii)
			for ivk, ivv := range flatten(prefix, iv) {
				res[ivk] = ivv
			}
		}

	case Map:
		for ik, iv := range vv {
			var prefix = propertyPath(key, ik)
			for ivk, ivv := range flatten(prefix, iv) {
				res[ivk] = ivv
			}
		}
	case map[string]interface{}:
		for ik, iv := range vv {
			var prefix = propertyPath(key, ik)
			for ivk, ivv := range flatten(prefix, iv) {
				res[ivk] = ivv
			}
		}
	default:
		res[key] = string(marshal(vv))
	}

	return res
}
