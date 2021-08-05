// @Author : a1234
// @Time : 2021/8/3 8:07 下午 
// @Describe :  

package simpleString

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	s := "anagram"
	ts := "nagaram"

	res := IsAnagram(s, ts)
	fmt.Println("res: ", res)
}

func TestIsAnagram1(t *testing.T) {
	s := "anagram"
	ts := "nagaram"

	res := IsAnagram1(s, ts)
	fmt.Println("res: ", res)
}