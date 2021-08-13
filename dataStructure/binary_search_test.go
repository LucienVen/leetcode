// @Author : a1234
// @Time : 2021/8/12 5:01 下午 
// @Describe :  

package dataStructure

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1,5,9, 15, 81, 89, 123, 189, 333}
	res := BinarySearch(arr, 1, 0, len(arr)-1)
	fmt.Println("res:", res)
}
