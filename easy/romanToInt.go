package easy

// https://leetcode.cn/problems/roman-to-integer/submissions/641113463/
// 13. 罗马数字转整数

var mapping = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// 若存在小的数字在大的数字的左边的情况，根据规则需要减去小的数字。
// 对于这种情况，我们也可以将每个字符视作一个单独的值，若一个数字右侧的数字比它大，则将该数字的符号取反。
// 但是这个解法无法排除错误的罗马数字写法
func romanToInt(s string) int {
	sum := 0
	for i := range s {
		cur := mapping[s[i]]
		if i == len(s)-1 {
			sum += cur
			return sum
		}
		next := mapping[s[i+1]]
		if cur < next {
			sum -= cur
		} else {
			sum += cur
		}
	}
	return sum
}
