package easy

import "testing"

// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
// 28. 找出字符串中第一个匹配项的下标

func Test_strStr(t *testing.T) {
	strStr("leetcode", "leeto")
}

func strStr(haystack string, needle string) int {
	hLen := len(haystack)
	nLen := len(needle)
	if hLen < nLen {
		return -1
	}
	maxIndex := hLen - nLen
	for i := 0; i <= maxIndex; i++ {
		if haystack[i:i+nLen] == needle {
			return i
		}
	}
	return -1
}
