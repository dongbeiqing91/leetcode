package easy

// https://leetcode.cn/problems/ransom-note/description/

func canConstruct(ransomNote string, magazine string) bool {
	list := [26]int{}
	for _, ch := range magazine {
		list[ch-'a']++
	}
	for _, ch := range ransomNote {
		count := list[ch-'a']
		if count == 0 {
			return false
		}
		list[ch-'a'] = count - 1
	}
	return true
}
