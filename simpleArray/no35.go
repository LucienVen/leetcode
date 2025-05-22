package simpleArray

// https://leetcode.cn/problems/search-insert-position/description/
// 35. 搜索插入位置

// tip: 相当于找第一个大于等于目标的数

func SearchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
	
	// 最后为什么返回的是left呢，因为一旦找不到target，
	//那么left一定会有等于right的一天(要么left变大，要么right在变小，无论循环怎么走，
	//都是有left或者right变化)，那么仅考虑最后left等于right时候，如果这个target是大于这个当前值，
	//那么left+1而我们的target也应该插入这个当前值的后面，所以说是left了，如果target小于当前值，
	//right-1，我们target就是应该插入当前值的位置，所以left没动，还是left

}
