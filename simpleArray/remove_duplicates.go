package simpleArray

import "fmt"

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var left = 0
	for right := 0; right < len(nums); right++ {
		if nums[left] != nums[right] {
			left += 1
			nums[left], nums[right] = nums[right], nums[left]
		}
	}

	nums = nums[:left+1]
	fmt.Println("nums:", nums)
	return len(nums)
}