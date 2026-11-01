package tryagain

import "strings"

func numDistinctIslands(grid [][]int) int {
	island := make(map[string]bool)
	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[i]); j++ {
			if grid[i][j] == 1 {
				sBuider := &strings.Builder{}
				dfs2(grid, i, j, sBuider, "z")
				island[sBuider.String()] = true
			}
		}
	}
	return len(island)
}

func dfs2(grid [][]int, i, j int, sBuilder *strings.Builder, direction string) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	sBuilder.WriteString(direction+",")
	dfs2(grid, i-1, j, sBuilder, "a")
	dfs2(grid, i, j-1, sBuilder, "b")
	dfs2(grid, i+1, j, sBuilder, "c")
	dfs2(grid, i, j+1, sBuilder, "d")
	sBuilder.WriteString("-"+direction+",")
}