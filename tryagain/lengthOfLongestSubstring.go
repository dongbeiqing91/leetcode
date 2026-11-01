package tryagain

func lengthOfLongestSubstring(s string) int {
	m := make(map[uint8]int)
	start := 0
	maxLen := 0
	for end := 0; end < len(s); end++ {
		if prev, ok := m[s[end]]; ok && prev >= start {
			start = prev + 1
		}
		m[s[end]] = end
		curLen := end - start + 1
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}
