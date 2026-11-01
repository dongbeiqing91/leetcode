package tryagain

func longestPalindrome(s string) string {
	var start int
	var maxLen int
	expand := func(left, right int) {
		for left >=0 && right < len(s) && s[left] == s[right] {
			if curLen := right - left + 1; curLen > maxLen {
				maxLen = curLen
			}
			left--
			right++
		}
	}
	for i := 0; i < len(s); i++ {
		expand(i, i)
		expand(i, i+1)
	}
	return s[start: start+maxLen]
}