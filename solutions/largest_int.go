package leetcode

import "math"

func Solution(A []int) int {
	// Implement your solution here
	maxValue := 0
	for i := 0; i < len(A); i++ {
		if math.Abs(float64(A[i])) > float64(maxValue) && contains(A, -A[i]) {
			maxValue = A[i]
		}
	}
	return maxValue
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
