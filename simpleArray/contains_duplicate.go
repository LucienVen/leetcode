// @Author : a1234
// @Time : 2021/7/23 11:28 上午 
// @Describe :  存在重复元素

// 输入: [1,2,3,4]
// 输出: false
// 输入: [1,1,1,3,3,4,3,2,4,2]
// 输出: true


package simpleArray

func ContainsDuplicate(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1{
		return false
	}

	// 使用map来实现set功能
	set := make(map[int]struct{})

	for _, num := range nums {
		set[num] = struct{}{}
	}

	if len(set) != len(nums) {
		return true
	}

	return false
}