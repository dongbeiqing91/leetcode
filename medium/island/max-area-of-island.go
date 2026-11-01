package island

// https://leetcode.cn/problems/max-area-of-island/

func maxAreaOfIsland(grid [][]int) int {
	var maxArea int
	row := len(grid)
	col := len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				area := flood(grid, i, j)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func flood(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return flood(grid, i+1, j) +
		flood(grid, i, j+1) +
		flood(grid, i-1, j) +
		flood(grid, i, j-1) + 1
}
