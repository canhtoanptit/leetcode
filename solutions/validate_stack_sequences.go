package leetcode

func validateStackSequences(pushed []int, popped []int) bool {
	i, j := 0, 0
	for _, e := range pushed {
		pushed[i] = e
		i++
		for i > 0 && pushed[i-1] == popped[j] {
			i--
			j++
		}
	}
	return i == 0
}
