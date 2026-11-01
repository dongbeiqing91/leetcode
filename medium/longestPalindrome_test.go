package medium

import (
	"log"
	"testing"
)

// https://leetcode.cn/problems/longest-palindromic-substring/
// 5. 最长回文子串

func longestPalindrome(s string) string {
	start, maxLen := 0, 0

	expand := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLen {
				start = left
				maxLen = right - left + 1
			}
			left--
			right++
		}
	}

	for i := 0; i < len(s); i++ {
		expand(i, i)   // 回文子串字符数为奇数，如bab
		expand(i, i+1) // 回文子串字符数为偶数，如abba
	}

	return s[start : start+maxLen]
}

func TestResult3(t *testing.T) {
	res := longestPalindrome("bab")
	log.Print(res)
	res2 := longestPalindrome("cbaab")
	log.Print(res2)
}
