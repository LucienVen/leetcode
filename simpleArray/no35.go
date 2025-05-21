package simpleArray

// https://leetcode.cn/problems/search-insert-position/description/
// 35. 搜索插入位置

func SearchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	// 二分查找
	// 如果相等，则返回下标
	// 如果大于，则返回下标
	// 如果小于，则返回下标

	mid := len(nums) / 2
	if nums[mid] == target {
		return mid
	} else if target < nums[mid] {
		// 小于中间值
		for i := mid; i >= 0; i-- {
			if nums[i] < target {
				return i + 1
			}
		}
		return 0
	} else if target > nums[mid] {
		// 大于中间值
		for i := mid; i < len(nums); i++ {
			if nums[i] > target {
				return i
			}
		}
		return len(nums)
	}

	return -1
}
