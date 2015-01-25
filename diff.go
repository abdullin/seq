package seq

import "strings"

func hasNestedObject(actual map[string]string, key string) bool {
	for k, _ := range actual {
		if strings.HasPrefix(k, key) {
			return true
		}
	}
	return false
}

func diff(expected, actual map[string]string) *Result {
	res := NewResult()

	for ek, ev := range expected {
		var av, ok = actual[ek]

		if !ok {
			if hasNestedObject(actual, ek) {
				res.AddDiff(ek, ev, "{Object}")
			} else {
				res.AddDiff(ek, ev, "nothing")
			}

		} else if av != ev {
			res.AddDiff(ek, ev, av)
		}
	}

	return res
}
