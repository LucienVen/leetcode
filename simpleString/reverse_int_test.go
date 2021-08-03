// @Author : a1234
// @Time : 2021/8/2 4:57 下午 
// @Describe :  

package simpleString

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	x := 1563847412
	res := Reverse(x)
	fmt.Println(res)
}

func TestReverse2(t *testing.T) {
	Reverse2(1)
}