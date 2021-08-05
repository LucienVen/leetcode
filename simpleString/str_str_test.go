// @Author : a1234
// @Time : 2021/8/5 7:11 下午 
// @Describe :  

package simpleString

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	haystack := "mississippi"
	needle := "issipi"


	res := StrStr(haystack, needle)
	fmt.Println("res:", res)
}