package easy

import (
	"log"
	"testing"
)

// https://leetcode.cn/problems/longest-common-prefix/
// 14. 最长公共前缀

func longestCommonPrefix(strs []string) string {
	s := strs[0]
    if s == "" {
        return ""
    }
    for i:=0;i<len(s);i++{
        for j:=1;j<len(strs);j++ {
            if i >= len(strs[j]) || strs[j][i] != s[i] {
                return s[0:i]
            }
        }
    }
    return s
}

func TestAA(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	prefix := longestCommonPrefix(strs)
	log.Println(prefix)
}
