// @Author : liangliangtoo
// @Time : 2021/7/28 3:35 下午 
// @Describe :  排序算法

package simpleArray

// 冒泡排序
func BubbleSort(old []int) []int {
	if len(old) == 0 || len(old) == 1 {
		return old
	}

	// 倒叙排序
	llen := len(old)

	for i := llen-1; i >= 0 ; i-- {
		for j := 0; j < i; j++ {
			if old[j] > old[i] {
				old[j], old[i] = old[i], old[j]
			}
		}
	}

	return old
}