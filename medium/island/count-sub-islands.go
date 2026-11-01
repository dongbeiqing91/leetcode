package island

// https://leetcode.cn/problems/count-sub-islands/

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
    row := len(grid1)
	col := len(grid1[0])
	for i:=0; i<row; i++ {
		for j:=0; j<col; j++ {
			if grid2[i][j] == 1 && grid1[i][j] == 0 {
				dfs4(grid2, i, j)
			}
		} 
	}
	var res int
	for i:=0; i<row; i++ {
		for j:=0; j<col; j++ {
			if grid2[i][j] == 1 {
				res ++
				dfs4(grid2, i, j)
			}
		} 
	}
	return res
}