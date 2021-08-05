// @Author : a1234
// @Time : 2021/8/3 8:31 下午 
// @Describe :  验证回文串

package simpleString

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	res := IsPalindrome(s)
	fmt.Println("res:", res)
}

func TestTs(t *testing.T) {
	Ts()
}