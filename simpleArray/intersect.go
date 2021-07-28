// @Author : liangliangtoo
// @Time : 2021/7/23 2:15 下午 
// @Describe :  给定两个数组，编写一个函数来计算它们的交集。

//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2,2]
//输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[4,9]

package simpleArray

func Intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	// 判断大小
	minNums, maxNums := nums1, nums2

	if len(nums1) > len(nums2) {
		minNums, maxNums = nums2, nums1
	}

	// 小数组在外层
	loop := make([]int, 0)
	//fmt.Println(minNums)
	//fmt.Println(maxNums)
	loopSet := make(map[int]int)
	for _, v1 := range minNums {
		for _, v2 := range maxNums {
			if v1 == v2 {
				if _, ok := loopSet[v1]; ok {
					loopSet[v1] += 1
				} else {
					loopSet[v1] = 1
				}

				//loop = append(loop, v1)
				break
			}
		}
	}


	for key, num := range loopSet {
		// 过滤数组
		for i := 1; i <= num; i++ {
			loop = append(loop, key)
		}
	}

	return loop
}

//func Intersect2(nums1 []int, nums2 []int) []int {
//	if len(nums1) == 0 || len(nums2) == 0 {
//		return nil
//	}
//	// 先排序
//	nums1 = BubbleSort(nums1)
//	nums2 = BubbleSort(nums2)
//
//	min := nums1
//	max := nums2
//	if len(nums1) > len(nums2) {
//		min, max = nums2, nums1
//	}
//
//
//
//
//}

func Intersect3(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	map1 := make(map[int]int)
	for _, value := range nums1 {
		if v, ok := map1[value]; !ok {
			map1[value] = 1
		} else {
			map1[value] = v + 1
		}
	}

	res := make([]int, 0)

	for _, item := range nums2 {
		if val, ok := map1[item]; ok {
			if val > 0 {
				map1[item] = val - 1
				res = append(res, item)
			} else {
				continue
			}
		}
	}

	return res

}