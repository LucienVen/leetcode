// @Author : a1234
// @Time : 2021/8/2 3:18 下午 
// @Describe :  

package simpleString

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	str := "hello"
	s := []byte(str)
	//fmt.Println("s: ", s)
	ReverseString(s)
}




func TestReverseString3(t *testing.T) {
	str := "hello"
	s := []byte(str)
	//fmt.Println("s: ", s)
	ReverseString3(s)
}