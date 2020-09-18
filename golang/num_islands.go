package main 

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	numIslands := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			numIslands := dfs(grid, i, j)
		}
	}

	return numIslands
}

func dfs(grid [][]byte, row, col int) int {
	if row >= len(grid) || row < 0 || col >= len(grid[0]) || col < 0 || grid[row][col] == '0' {
		return 0
	}

	grid[row][col] = '0'
	dfs(grid, row+1, col)
	dfs(grid, row-1, col)
	dfs(grid, row, col + 1)
	dfs(grid, row, col - 1)
    return 1
}