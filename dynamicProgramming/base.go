// @Author : a1234
// @Time : 2021/8/5 10:21 下午 
// @Describe :  

package dynamicProgramming

var SolutionMap map[int]int

func init()  {
	SolutionMap = make(map[int]int)
}


func Solution(n int32) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return Solution(n - 1) + Solution(n - 2)
}


func Solution2(n int32) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return 0
}