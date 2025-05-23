package simpleArray

import "testing"

func TestMerge(t *testing.T) {

	cases := []struct {
		Name     string
		Input    [][]int
		Expected [][]int
	}{
		{
			Name:     "test1",
			Input:    [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			Expected: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			Name:     "test2",
			Input:    [][]int{{1, 4}, {2, 3}},
			Expected: [][]int{{1, 4}},
		},
		{
			Name:     "test3",
			Input:    [][]int{{1, 4}, {0, 2}, {3, 5}},
			Expected: [][]int{{0, 5}},
		},
		{
			Name:     "test4",
			Input:    [][]int{{1, 4}, {2, 3}, {5, 7}, {6, 8}},
			Expected: [][]int{{1, 4}, {5, 8}},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			result := Merge(c.Input)
			if len(result) != len(c.Expected) {
				t.Log("Input: ", c.Input)
				t.Errorf("expected %v, got %v", c.Expected, result)
				return
			}
			for i := range result {
				if result[i][0] != c.Expected[i][0] || result[i][1] != c.Expected[i][1] {
					t.Log("Input: ", c.Input)
					t.Errorf("expected %v, got %v", c.Expected, result)
					break
				}
			}
		})
	}

}

func TestMerge2(t *testing.T) {

	cases := []struct {
		Name     string
		Input    [][]int
		Expected [][]int
	}{
		{
			Name:     "test1",
			Input:    [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			Expected: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			Name:     "test2",
			Input:    [][]int{{1, 4}, {2, 3}},
			Expected: [][]int{{1, 4}},
		},
		{
			Name:     "test3",
			Input:    [][]int{{1, 4}, {0, 2}, {3, 5}},
			Expected: [][]int{{0, 5}},
		},
		{
			Name:     "test4",
			Input:    [][]int{{1, 4}, {2, 3}, {5, 7}, {6, 8}},
			Expected: [][]int{{1, 4}, {5, 8}},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			result := Merge2(c.Input)
			if len(result) != len(c.Expected) {
				t.Log("Input: ", c.Input)
				t.Errorf("expected %v, got %v", c.Expected, result)
				return
			}
			for i := range result {
				if result[i][0] != c.Expected[i][0] || result[i][1] != c.Expected[i][1] {
					t.Log("Input: ", c.Input)
					t.Errorf("expected %v, got %v", c.Expected, result)
					break
				}
			}
		})
	}

}
