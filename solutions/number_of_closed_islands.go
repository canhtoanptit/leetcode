package leetcode

func closedIsland(grid [][]int) (res int) {
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] == 0 && flood(grid, i, j) {
				res++
			}
		}
	}

	return res
}

func flood(grid [][]int, i, j int) bool {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
		return false
	}

	if grid[i][j] == 1 {
		return true
	}

	grid[i][j] = 1

	up := flood(grid, i-1, j)
	down := flood(grid, i+1, j)
	left := flood(grid, i, j-1)
	right := flood(grid, i, j+1)

	return up && down && left && right
}
