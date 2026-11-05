package easy

// https://leetcode.cn/problems/pascals-triangle/

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for row := range res {
		res[row] = make([]int, row+1)
		res[row][0] = 1
		res[row][row] = 1
		for i := 1; i < row; i++ {
			res[row][i] = res[row-1][i-1] + res[row-1][i]
		}
	}
	return res
}
