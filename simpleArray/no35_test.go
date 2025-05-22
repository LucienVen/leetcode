package simpleArray

import "testing"

// no. 35
func TestSearchInsert(t *testing.T) {
	cases := []struct {
		Name     string
		Input    []int
		Target   int
		Expected int
	}{{
		Name:     "test1",
		Input:    []int{1, 3, 5, 6},
		Target:   5,
		Expected: 2,
	}, {
		Name:     "test2",
		Input:    []int{1, 3, 5, 6},
		Target:   2,
		Expected: 1,
	}, {
		Name:     "test3",
		Input:    []int{1, 3, 5, 6},
		Target:   7,
		Expected: 4,
	}, {
		Name:     "test4",
		Input:    []int{1, 3},
		Target:   4,
		Expected: 2,
	},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			result := SearchInsert(c.Input, c.Target)
			if result != c.Expected {
				t.Errorf("expected %d, got %d", c.Expected, result)
			}

			//t.Log(result)
		})

		//result := SearchInsert(c.Input, c.Target)
		//t.Log("Name: ", c.Name, "result: ", result)
	}

	//t.Log("no. 35 test passed")
}
