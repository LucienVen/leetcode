// @Author : a1234
// @Time : 2021/8/5 9:36 下午
// @Describe :

package simpleString

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	flag := strs[0] // 设第一个字符串的第一个字符为
	return flag
}
