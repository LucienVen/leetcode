// @Author : a1234
// @Time : 2021/7/28 4:58 下午 
// @Describe :  给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

//输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]

package simpleArray

import "fmt"

func MoveZeroes(nums []int)  {
	if len(nums) == 0 {
		return
	}

	if len(nums) == 1 && nums[0] == 0 {
		return
	}

	llen := len(nums)
	index := 0

	for i := 0; i < llen; i++ {
		if nums[i] != 0 {

			nums[index] = nums[i]
			index = index + 1
		}
	}

	for index < llen {

		nums[index] = 0
		index = index + 1
	}

	fmt.Println(nums)

}


func MoveZeroes2(nums []int) {

}