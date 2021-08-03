// @Author : a1234
// @Time : 2021/8/2 4:57 下午 
// @Describe :  

package simpleString

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Reverse(x int) int {
	if x == 0 {
		return x
	}

	// 转换成字符串
	str := strconv.Itoa(x)
	fmt.Println("str:", str)

	//strByte := []byte(str)

	// 判断第一个字符，是否 - 号
	flag := strings.Contains(str, "-")

	if flag {
		str = str[1:]
	}

	//fmt.Println(str)
	var (
		left = 0
		right = len(str) - 1
	)

	strByte := []byte(str)

	for left < right {
		strByte[left], strByte[right] = strByte[right], strByte[left]
		left += 1
		right -= 1
	}

	strReverse := string(strByte)
	//fmt.Println(strReverse)


	// string to int
	res, _ := strconv.Atoi(strReverse)

	//fmt.Println(2<<32)
	//fmt.Println(-(2<<32))

	// 判断是否超出32位
	if res < -(2<<32) || res > 2<<32 {
		return 0
	}	

	if flag {
		res = res * -1
	}



	return res
}


// 反转整数
func Reverse2(x int) int {
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MinInt32)
	return 0
}