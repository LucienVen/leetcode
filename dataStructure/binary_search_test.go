// @Author : a1234
// @Time : 2021/8/12 5:01 下午
// @Describe :

package dataStructure

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		Name   string
		Arr    []int
		Target int
		Expect int
	}{
		{
			Name:   "test1",
			Arr:    []int{1, 5, 9, 15, 81, 89, 123, 189, 333},
			Target: 1,
			Expect: 0,
		}, {
			Name:   "test2",
			Arr:    []int{1, 3, 5, 6},
			Target: 5,
			Expect: 2,
		},
	}

	for _, c := range cases {

		t.Run(c.Name, func(t *testing.T) {
			result := BinarySearch(c.Arr, c.Target, 0, len(c.Arr)-1)
			if result != c.Expect {
				t.Errorf("expected %d, got %d", c.Expect, result)
			}
		})
	}

	t.Log("Binary search test passed")

	//arr := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
	//res := BinarySearch(arr, 1, 0, len(arr)-1)
	//fmt.Println("res:", res)
}
