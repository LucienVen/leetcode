// @Author : a1234
// @Time : 2021/7/28 3:44 下午 
// @Describe :  

package simpleArray

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	todo := []int{2,3,6,1,10,28,6}
	res := BubbleSort(todo)
	fmt.Println(res)
}
