// @Author : a1234
// @Time : 2021/8/2 3:17 下午 
// @Describe :  反转字符串

//输入：["h","e","l","l","o"]
//输出：["o","l","l","e","h"]

//输入：["H","a","n","n","a","h"]
//输出：["h","a","n","n","a","H"]

package simpleString

import "fmt"

func ReverseString(s []byte)  {
	if len(s) == 0 || len(s) == 1 {
		return
	}

	llen := len(s)

	for i := 0; i < llen/2; i++ {
		s[i], s[llen-1-i] = s[llen-1-i], s[i]
	}

	fmt.Println(string(s))
}

// 递归
func ReverseString2(s []byte) {
	if len(s) == 0 || len(s) == 1 {
		return
	}


}

// 双指针
func ReverseString3(s []byte) {
	if len(s) == 0 || len(s) == 1 {
		return
	}

	var (
		left = 0
		right = len(s) - 1
	)

	for left < right {
		s[left], s[right] = s[right], s[left]
		left += 1
		right -= 1
	}

	fmt.Println(string(s))

}