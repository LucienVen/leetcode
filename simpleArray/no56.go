package simpleArray

import (
	"sort"
)

// 56. 合并区间
// https://leetcode.cn/problems/merge-intervals/description/

func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals); i++ {
		handle := intervals[i]

		for j := i + 1; j < len(intervals); j++ {
			target := intervals[j]

			merge := max(handle[0], target[0]) <= min(handle[1], target[1])
			if merge {
				// 区间重叠
				// 移除元素
				intervals = append(intervals[:j], intervals[j+1:]...)
				intervals = append(intervals[:i], intervals[i+1:]...)

				intervals = append(intervals, []int{min(handle[0], target[0]), max(target[1], handle[1])})

				i = -1 // 重新再来
				break
			}
		}
	}

	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	return intervals
}

func Merge2(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	start, end := intervals[0][0], intervals[0][1]
	ans := make([][]int, 0)

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > end {
			// 新的元素，左端点大于当前的右端点，则表示不连续

			// 旧元素存入结果
			ans = append(ans, []int{start, end})
			// 新区间重新开始
			start, end = intervals[i][0], intervals[i][1]
		} else {
			// 重叠
			end = max(end, intervals[i][1])
		}
	}

	// 最后区间存入
	ans = append(ans, []int{start, end})

	return ans
}
