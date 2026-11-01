package island

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/number-of-distinct-islands/description/

func numDistinctIslands(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 记录所有岛屿的序列化结果
	islands := make(map[string]bool)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				// 淹掉这个岛屿，同时存储岛屿的序列化结果
				var sb strings.Builder
				// 初始的方向可以随便写，不影响正确性
				dfs2(grid, i, j, &sb, 666)
				islands[sb.String()] = true
			}
		}
	}
	// 不相同的岛屿数量
	return len(islands)
}

func dfs2(grid [][]int, i int, j int, sb *strings.Builder, dir int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 0 {
		return
	}
	// 前序遍历位置：进入 (i, j)
	grid[i][j] = 0
	sb.WriteString(strconv.Itoa(dir) + ",")

	// 搜索相邻节点
	dfs2(grid, i-1, j, sb, 1)
	dfs2(grid, i+1, j, sb, 2)
	dfs2(grid, i, j-1, sb, 3)
	dfs2(grid, i, j+1, sb, 4)

	// 后序遍历位置：离开 (i, j)
	sb.WriteString(strconv.Itoa(-dir) + ",")
}
