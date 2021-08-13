// @Author : liangliangtoo
// @Time : 2021/8/12 4:36 下午 
// @Describe :  二分查找法

package dataStructure


// 二分查找法 - 递归
func BinarySearch(array []int, target int, l, r int) int {
	if l > r {
		// 出界
		return -1
	}

	mid := (l + r) / 2
	middleNum := array[mid]

	if middleNum == target {
		return mid
	} else if middleNum > target {
		return BinarySearch(array, target, l, mid)
	} else if middleNum < target {
		return BinarySearch(array, target, mid, r)
	}

	return -1
}