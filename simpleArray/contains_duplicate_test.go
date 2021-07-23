// @Author : a1234
// @Time : 2021/7/23 11:34 上午 
// @Describe :  

package simpleArray

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	nums := []int{1,1,1,3,3,4,3,2,4,2}
	res := ContainsDuplicate(nums)
	fmt.Println("result:", res)
}
