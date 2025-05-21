package simpleArray

import "testing"

func TestPivotIndex(t *testing.T) {
	cases := []struct {
		Name     string
		Input    []int
		Expected int
	}{
		{
			Name:     "test1",
			Input:    []int{1, 7, 3, 6, 5, 6},
			Expected: 3,
		},
		{
			Name:     "test2",
			Input:    []int{1, 2, 3},
			Expected: -1,
		},
		{
			Name:     "test3",
			Input:    []int{2, 1, -1},
			Expected: 0,
		},
		{
			Name:     "test4",
			Input:    []int{1, 2, 3, 4, 5},
			Expected: -1,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			result := PivotIndex(c.Input)
			if result != c.Expected {
				t.Errorf("expected %d, got %d", c.Expected, result)
			}
		})
	}
}
