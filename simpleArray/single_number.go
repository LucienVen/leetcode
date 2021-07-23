// @Author : a1234
// @Time : 2021/7/23 12:00 下午 
// @Describe :  只出现一次的数字

//输入: [4,1,2,1,2]
//输出: 4
//输入: [2,2,1]
//输出: 1

package simpleArray


func SingleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var reduce = 0
	for _, num := range nums {
		reduce = reduce ^ num
	}

	return reduce
}

