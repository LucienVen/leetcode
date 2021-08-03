// @Author : a1234
// @Time : 2021/8/3 10:47 上午 
// @Describe :  字符串中的第一个唯一字符

//s = "leetcode"
//返回 0

//s = "loveleetcode"
//返回 2

package simpleString

func FirstUniqChar(s string) int {

	numMap := make(map[string]int)
	for _, v := range s {
		if _, ok := numMap[string(v)]; ok {
			numMap[string(v)] = numMap[string(v)] + 1
		} else {
			numMap[string(v)] = 1
		}
	}


	elem := -1
	for index, item := range s {
		if value, ok := numMap[string(item)]; ok {
			if value == 1 {
				elem = index
				break
			}
		}
	}

	return elem
}
