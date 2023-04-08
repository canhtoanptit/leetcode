package leetcode

func numEnclaves(grid [][]int) int {
	var result int
	for i := 1; i < len(grid)-1; i++ {
		for j := 0; j < len(grid)-1; j++ {
			if grid[i][j] == 1 && check(grid, i, j) {
				result++
			}
		}
	}
	return result
}

func check(grid [][]int, i, j int) bool {
	if i < 0 || j < 0 || i > len(grid) || j > len(grid) {
		return false
	}

	if grid[i][j] == 0 {
		return true
	}

	grid[i][j] = 0

	up := check(grid, i-1, j)
	down := check(grid, i+1, j)
	left := check(grid, i, j-1)
	right := check(grid, i, j+1)

	return up && down && left && right
}
