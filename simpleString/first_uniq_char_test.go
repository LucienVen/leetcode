// @Author : a1234
// @Time : 2021/8/3 10:47 上午 
// @Describe :  

package simpleString

import (
	"fmt"
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	target := "leetcode"
	res := FirstUniqChar(target)
	fmt.Println("index: ", res)
}