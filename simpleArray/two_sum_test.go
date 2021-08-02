// @Author : a1234
// @Time : 2021/8/2 1:55 下午 
// @Describe :  

package simpleArray

import (
	"fmt"
	"testing"
)

var (
	tempNums = []int{11, 15, 2, 7}
	tempTarge = 9
)

func TestTwoSum1(t *testing.T) {
	nums := []int{2,7,11,15}
	target := 9
	res := TwoSum1(nums, target)
	fmt.Println(res)
}

func TestTwoSum2(t *testing.T) {
	nums := []int{2,7,11,15}
	target := 9
	res := TwoSum2(nums, target)
	fmt.Println(res)
}


// 性能测试
func BenchmarkTwoSum1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum1(tempNums, tempTarge)
	}
}

func BenchmarkTwoSum2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum2(tempNums, tempTarge)
	}
}