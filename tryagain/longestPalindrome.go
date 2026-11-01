package tryagain

func longestPalindrome(s string) string {
	var maxLen int
	var start int
	expand := func(left, right int) {
		for left >=0 && right <= len(s) - 1 {
			if s[left] == s[right] {
				if len := right - left + 1; len > maxLen {
					start = left
					maxLen = len
				}
				left--
				right++
			} else {
				break
			}
		}
	}
	for i := 0; i < len(s); i++ {
		expand(i, i)
		expand(i, i+1)
	}
	return s[start : start+maxLen]
}
