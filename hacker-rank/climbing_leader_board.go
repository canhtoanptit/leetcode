package hacker_rank

import "fmt"

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	var rs []int32

	var no_dup []int32

	// map score and index
	mm := make(map[int32]int)
	tem := 0
	for _, e := range ranked {
		_, ok := mm[e]
		if !ok {
			tem++
			mm[e] = tem
			no_dup = append(no_dup, e)
		}
	}

	for _, p := range player {
		fmt.Println(p)
	}
	return rs
}
