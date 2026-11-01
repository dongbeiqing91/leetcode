package island

// https://leetcode.cn/problems/number-of-closed-islands/
func closedIsland(grid [][]int) int {
	res := 0
	row := len(grid)
	col := len(grid[0])
	for i := 0; i < row; i++ {
		dfs3(grid, i, 0)
		dfs3(grid, i, col-1)
	}
	for j := 0; j < col; j++ {
		dfs3(grid, 0, j)
		dfs3(grid, row-1, j)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 0 {
				res++
				dfs3(grid, i, j)
			}
		}
	}
	return res
}

func dfs3(grid [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
		return
	}
	if grid[i][j] == 1 {
		return
	}
	grid[i][j] = 1
	dfs3(grid, i-1, j)
	dfs3(grid, i, j-1)
	dfs3(grid, i+1, j)
	dfs3(grid, i, j+1)
}
