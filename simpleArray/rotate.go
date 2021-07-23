// @Author : a1234
// @Time : 2021/7/22 6:14 下午 
// @Describe :  旋转数组

package simpleArray

import "fmt"

//输入: nums = [1,2,3,4,5,6,7], k = 3
//输出: [5,6,7,1,2,3,4]

func Rotate(nums []int, k int)  {
	if len(nums) == 0 {
		return
	}

	if k > len(nums) {
		k = k%len(nums)
	}

	temp := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		index := (i+len(nums)-k)%len(nums) // 余数计算
		//fmt.Println("index", index)
		//nums[i], nums[index] = nums[index], nums[i]
		temp = append(temp, nums[index])
	}

	for index, value := range temp {
		nums[index] = value
	}



	fmt.Println(nums)
}