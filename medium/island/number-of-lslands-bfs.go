package island

// numIslandsBFS 使用广度优先搜索返回岛屿数量。
func numIslandsBFS(grid [][]byte) int {
	// 空网格没有岛屿。
	if len(grid) == 0 {
		return 0
	}

	rows, columns := len(grid), len(grid[0])
	// 第一行为空时没有可以遍历的列。
	if columns == 0 {
		return 0
	}

	count := 0
	// 定义上下左右四个搜索方向。
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if grid[i][j] != '1' {
				continue
			}

			// 遇到一块尚未访问的陆地，岛屿数量加一。
			count++
			grid[i][j] = '0'
			// 队列用于保存待访问的陆地坐标。
			queue := [][]int{{i, j}}

			for len(queue) > 0 {
				// 取出队首坐标，并从队列中移除。
				cell := queue[0]
				queue = queue[1:]

				for _, direction := range directions {
					nextRow := cell[0] + direction[0]
					nextColumn := cell[1] + direction[1]
					// 检查越界，并只把相邻的陆地加入队列。
					if nextRow < 0 || nextRow >= rows ||
						nextColumn < 0 || nextColumn >= columns ||
						grid[nextRow][nextColumn] != '1' {
						continue
					}

					// 入队前标记为已访问，避免重复加入队列。
					grid[nextRow][nextColumn] = '0'
					queue = append(queue, []int{nextRow, nextColumn})
				}
			}
		}
	}

	return count
}
