package medium

import (
	"log"
	"testing"
)

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
// 3. 无重复字符的最长子串

func lengthOfLongestSubstring(s string) int {
	// last 记录每个字符最近一次出现的位置。
	last := make(map[byte]int)

	// start 是当前无重复字符窗口的左边界。
	start := 0
	maxLen := 0

	// end 是当前窗口的右边界，用于逐字符向右扫描。
	for end := 0; end < len(s); end++ {
		// 如果当前字符之前出现过，并且上次位置仍在窗口内，
		// 说明窗口中产生了重复字符，将左边界移到该位置之后。
		if previous, ok := last[s[end]]; ok && previous >= start {
			start = previous + 1
		}

		// 更新当前字符最近一次出现的位置。
		last[s[end]] = end

		// 当前窗口为 [start, end]，计算长度并更新最长长度。
		if currentLen := end - start + 1; currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}

func TestResult(t *testing.T) {
	res := lengthOfLongestSubstring("tmmzuxt")
	log.Print(res)
}
