package easy

// https://leetcode.cn/problems/valid-parentheses/description/

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	var list []string
	m := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	for i := range s {
		char := string(s[i])
		if n, ok := m[char]; ok {
			if len(list) > 0 && list[len(list)-1] == n {
				list = list[0 : len(list)-1]
			} else {
				return false
			}
		} else {
			list = append(list, char)
		}
	}
	return len(list) == 0
}
