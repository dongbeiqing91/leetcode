package medium

import (
	"log"
	"testing"
)

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
// 3. 无重复字符的最长子串

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	maxLen := 1
	m := make(map[byte]int)
	m[s[0]] = 0
	startIndex := 0
	endIndex := 1
	currentLen := 1
	for index := 1; index < len(s); index++ {
		endIndex = index
		if foundIndex, found := m[s[index]]; found {
			oldStartIndex := startIndex
			startIndex = foundIndex + 1
			for i := oldStartIndex; i < startIndex; i++ {
				delete(m, s[i])
			}
		} else {
			currentLen = endIndex - startIndex + 1
		}
		m[s[index]] = index
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}

func TestResult(t *testing.T) {
	res := lengthOfLongestSubstring("tmmzuxt")
	log.Print(res)
}
