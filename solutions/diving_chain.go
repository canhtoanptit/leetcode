package leetcode

func SolutionDivingChain(A []int) int {
	return minCostToDivideChain(A)
}

func minCostToDivideChain(A []int) int {
	n := len(A)
	if n < 5 {
		return 0
	}

	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
		for j := range dp[i] {
			if i == j {
				dp[i][j] = 0
			} else if j-i >= 2 {
				dp[i][j] = 1e9
			}
		}
	}

	for i := 1; i < n-1; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if k-i <= 1 {
					continue
				}

				cost := A[i] + A[j] + A[k]
				cost += dp[i+1][j-1] + dp[i][j] + dp[j][k] + dp[k+1][n-2]
				dp[i][k] = min(dp[i][k], cost)
			}
		}
	}

	return dp[1][n-2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
