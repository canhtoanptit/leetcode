package hacker_rank

func nonDivisibleSubset(k int, s []int) int {
	// Write your code here
	var arr [][]int
	for i, e := range s {
		for j := i + 1; j < len(s); j++ {
			if (e+s[j])%k == 0 {
				arr = append(arr, []int{e, s[j]})
			}
		}
	}
	return 1
}
