// @Author : a1234
// @Time : 2021/7/28 7:01 下午 
// @Describe :  两数之和

// 输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
//

package simpleArray

// 暴力遍历
func TwoSum1(nums []int, target int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}

	llen := len(nums)

	for i := 0; i < llen; i++ {
		for j := i+1; j < llen; j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// hashmap - 补数
func TwoSum2(nums []int, target int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}

	resMap := make(map[int]int)

	for i, v := range nums {
		sep := target - v
		if j, ok := resMap[sep]; ok {
			return []int{i, j}
		} else {
			resMap[v] = i
		}

	}

	return []int{}
}