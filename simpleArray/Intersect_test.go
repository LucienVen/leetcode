// @Author : a1234
// @Time : 2021/7/23 2:22 下午 
// @Describe :  

package simpleArray

import (
	"fmt"
	"testing"
)

func TestIntersect(t *testing.T) {
	//nums1 := []int{1,2,2,1}
	//nums2 := []int{2,2}

	//nums1 := []int{3,1,2}
	//nums2 := []int{1,1}

	nums1 := []int{4,9,5}
	nums2 := []int{9,4,9,8,4}

	res := Intersect(nums1, nums2)
	fmt.Println(res)

}