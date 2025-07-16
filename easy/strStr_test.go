package easy

import "testing"

// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
// 28. 找出字符串中第一个匹配项的下标

func Test_strStr(t *testing.T) {
		strStr("leetcode", "leeto")
}

func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	haystackLen := len(haystack)
	if haystackLen < needleLen {
		return -1
	}
	maxIndex := haystackLen - needleLen
	index := 0
	for {
		if index >= haystackLen {
			return -1
		}
		for index < haystackLen && haystack[index] != needle[0] {
			index++
		}
		if index > maxIndex {
			return -1
		}
		if haystack[index:index+needleLen] == needle {
			return index
		}
		index++
	}
}
