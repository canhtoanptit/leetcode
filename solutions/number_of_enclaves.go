package leetcode

func numEnclaves(grid [][]int) (res int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i*j == 0 || i == len(grid)-1 || j == len(grid[i])-1 {
				if grid[i][j] == 1 {
					fill(grid, i, j)
				}
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				res++
			}
		}
	}

	return res
}

func fill(grid [][]int, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
		return
	}

	if grid[i][j] == 0 {
		return
	}

	grid[i][j] = 0

	fill(grid, i-1, j)
	fill(grid, i+1, j)
	fill(grid, i, j-1)
	fill(grid, i, j+1)
}
