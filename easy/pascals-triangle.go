package easy

// https://leetcode.cn/problems/pascals-triangle/

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for line := range res {
		res[line] = make([]int, line+1)
		res[line][0] = 1
		res[line][line] = 1
		for i := 1; i < line; i++ {
			res[line][i] = res[line-1][i-1] + res[line-1][i]
		}
	}
	return res
}
