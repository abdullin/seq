package seq

import "testing"

type Robot struct {
	Legs int    `json:"legs"`
	Arms int    `json:"arms"`
	Name string `json:"name"`
}

type Party struct {
	Rating  []int             `json:"rating"`
	Seating map[string]*Robot `json:"seating"`
}

func TestPartialSimpleObject(t *testing.T) {
	expect := Map{
		"legs": 2,
	}

	actual := Robot{
		Legs: 2,
		Arms: 2,
	}

	result := expect.Test(actual)
	expectOk(result, t)
}

func TestExactSimpleObject(t *testing.T) {
	expect := Map{
		"legs": 2,
		"arms": 2,
		"name": "benny",
	}

	actual := Robot{
		Legs: 2,
		Arms: 2,
		Name: "benny",
	}

	result := expect.Test(actual)
	expectOk(result, t)
}

func TestWrongSimpleObject(t *testing.T) {
	expect := Map{
		"legs": 2,
		"arms": 2,
		"name": "benny",
	}

	actual := Robot{
		Legs: 2,
		Arms: 2,
		Name: "Bender",
	}

	result := expect.Test(actual)
	expectFail(result, t)
}

func TestSimpleArray(t *testing.T) {
	expect := Map{
		"[0]": Map{
			"name": "R2D2",
		},
		"[1]": Map{
			"name": "C3PO",
		},
	}

	result := expect.Test([]*Robot{
		&Robot{Name: "R2D2"},
		&Robot{Name: "C3PO"},
	})
	expectOk(result, t)
}

func TestExactComplexObject(t *testing.T) {
	actual := &Party{
		Rating: []int{4, 5, 4},
		Seating: map[string]*Robot{
			"front": &Robot{
				Name: "R2D2",
				Arms: 1,
				Legs: 3,
			},
			"back": &Robot{
				Name: "Marvin",
				Legs: 2,
				Arms: 2,
			},
			"right": &Robot{
				Name: "C3PO",
				Legs: 2,
				Arms: 2,
			},
		},
	}

	expect := Map{
		// array
		"rating.len": 3,
		"rating[1]":  5,
		// flat path with value terminator
		"seating.front.name": "R2D2",
		"seating.front.arms": "1",
		"seating.front.legs": 3,
		// flat path with map terminator
		"seating.right": Map{
			"name": "C3PO",
		},
		// flat path with object terminator
		"seating.back": &Robot{
			Name: "Marvin",
			Legs: 2,
			Arms: 2,
		},
	}
	result := expect.Test(actual)
	expectOk(result, t)
}

func expectOk(result *Result, t *testing.T) {
	if result.Ok() {
		return
	}
	t.Error("Test should pass")
	for i, l := range result.Diffs {
		t.Log(i, l)
	}
}

func expectFail(result *Result, t *testing.T) {

	if !result.Ok() {
		return
	}
	t.Error("Test should fail")
	for i, l := range result.Diffs {
		t.Log(i, l)
	}
}