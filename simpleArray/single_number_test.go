// @Author : a1234
// @Time : 2021/7/23 12:00 下午 
// @Describe :  

package simpleArray

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	nums := []int{4,1,2,1,2}
	res := SingleNumber(nums)
	fmt.Println(res)
}