// @Author : a1234
// @Time : 2021/8/3 8:31 下午 
// @Describe :  验证回文串

package simpleString

import (
	"fmt"
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}

	if len(s) == 1 {
		return true
	}

	s = strings.ToLower(s)
	fmt.Println(s)

	left := 0
	right := len(s) - 1

	for left < right {
		if !unicode.IsDigit(rune(s[left])) && !unicode.IsLetter(rune(s[left])) {
			left += 1
			continue
		}

		if !unicode.IsDigit(rune(s[right])) && !unicode.IsLetter(rune(s[right])) {
			right -= 1
			continue
		}

		fmt.Println(fmt.Sprintf("left: %s, right: %s", string(s[left]), string(s[right])))

		if s[left] != s[right] {
			return false
		}

		left += 1
		right -= 1
	}

	return true
}

func Ts()  {
	fmt.Println((17638923.2)/17639323.2)
	 res := fmt.Sprintf("%.2f", 17638923.2*100/17639323.2*100) + "%"  // 毛利率=毛利润/总充值金额
	 fmt.Println(res)

	 fmt.Println(fmt.Sprintf("%.2f", 0.3546879765))
}

