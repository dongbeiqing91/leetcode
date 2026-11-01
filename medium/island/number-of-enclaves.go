package island

// https://leetcode.cn/problems/number-of-enclaves

func numEnclaves(grid [][]int) int {
	var res int
	row := len(grid)
	col := len(grid[0])
	for i := 0; i < row; i++ {
		dfs4(grid, i, 0)
		dfs4(grid, i, col-1)
	}
	for j := 0; j < col; j++ {
		dfs4(grid, 0, j)
		dfs4(grid, row-1, j)
	}
	for i := 1; i < row-1; i++ {
		for j := 1; j < col-1; j++ {
			if grid[i][j] == 1 {
				res++
			}
		}
	}
	return res
}

func dfs4(grid [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	dfs4(grid, i-1, j)
	dfs4(grid, i, j-1)
	dfs4(grid, i+1, j)
	dfs4(grid, i, j+1)
}
