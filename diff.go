package seq

import (
	"fmt"
	"strings"
)

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
