package simpleArray

// https://leetcode.cn/problems/find-pivot-index/
// 724. 寻找数组的中心下标

func PivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	leftSum := 0
	for i, num := range nums {
		if leftSum == sum-leftSum-num {
			return i
		}
		leftSum += num
	}

	return -1
}
