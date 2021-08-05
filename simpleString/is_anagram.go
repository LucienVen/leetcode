// @Author : a1234
// @Time : 2021/8/3 8:07 下午 
// @Describe :  有效的字母异位词

package simpleString

import (
	"sort"
	"strings"
)

func IsAnagram(s string, t string) bool {
	if len(s) == 0 || len(t) == 0 {
		return false
	}

	if len(s) != len(t) {
		return false
	}

	smap := make(map[string]int)
	tmap := make(map[string]int)

	for _, v := range s {
		if val, ok := smap[string(v)]; ok {
			smap[string(v)] = val + 1
		} else {
			smap[string(v)] = 1
		}
	}

	for _, v := range t {
		if val, ok := tmap[string(v)]; ok {
			tmap[string(v)] = val + 1
		} else {
			tmap[string(v)] = 1
		}
	}

	if len(smap) != len(tmap) {
		return false
	}

	for k, v := range smap {
		num, ok := tmap[k];
		if !ok {
			return false
		}

		if num != v {
			return false
		}
	}

	return true
}

// 字符串 排序后是否相等
func IsAnagram1(s string, t string) bool {
	if len(s) == 0 || len(t) == 0 {
		return false
	}

	if len(s) != len(t) {
		return false
	}

	slist := strings.Split(s, "")
	sort.Strings(slist)
	ns := strings.Join(slist, "")

	//fmt.Println("ns:", ns)

	tlist := strings.Split(t, "")
	sort.Strings(tlist)
	nt := strings.Join(tlist, "")
	//fmt.Println("nt:", nt)

	if ns == nt {
		return true
	}

	return false


}

