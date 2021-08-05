// @Author : a1234
// @Time : 2021/8/5 7:11 下午 
// @Describe :  给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。

//输入：haystack = "hello", needle = "ll"
//输出：2

package simpleString

func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		 return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	if len(needle) == len(haystack) {
		if needle != haystack {
			 return -1
		}
	}

	flag := needle[0]

	for index, letter := range haystack {
		if letter == int32(flag) {
			right := index+len(needle)
			if right > len(haystack) {
				continue
			}
			// 开始截断
			//fmt.Println("index:", index, " | index+len(needle):", index+len(needle))
			//fmt.Println("haystack[index:index+len(needle)]", haystack[index:index+len(needle)])



			if haystack[index:index+len(needle)] == needle {
				return index
			}
		}

	}

	return -1
}