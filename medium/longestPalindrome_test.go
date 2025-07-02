package medium

import (
	"log"
	"testing"
)

// https://leetcode.cn/problems/longest-palindromic-substring/
// 5. 最长回文子串

func longestPalindrome(s string) string {
	maxLen := 1
	startIndex := 0

	for i := 0; i < len(s); i++ {
		// 回文子串字符数为奇数，如bab
		l := i
		r := i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			length := r - l + 1
			if length > maxLen {
				maxLen = length
				startIndex = l
			}
			l--
			r++
		}

		// 回文子串字符数为偶数，如abba
		l = i
		r = i + 1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			length := r - l + 1
			if length > maxLen {
				maxLen = length
				startIndex = l
			}
			l--
			r++
		}
	}

	return s[startIndex : startIndex+maxLen]
}

func TestResult3(t *testing.T) {
	res := longestPalindrome("bab")
	log.Print(res)
	res2 := longestPalindrome("cbaab")
	log.Print(res2)
}
