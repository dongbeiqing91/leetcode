package easy

import (
	"log"
	"testing"
)

// https://leetcode.cn/problems/longest-common-prefix/
// 14. 最长公共前缀

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	var prefix []byte
	for i := 0; ; i++ {
		var compare byte
		for index := range strs {
			str := strs[index]
			if i > len(str)-1 {
				return string(prefix)
			}
			cur := str[i]
			if index == 0 {
				compare = cur
				continue
			}
			if cur != compare {
				return string(prefix)
			} else if index == len(strs)-1 {
				prefix = append(prefix, cur)
			}
		}
	}
}

func TestAA(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	prefix := longestCommonPrefix(strs)
	log.Println(prefix)
}
