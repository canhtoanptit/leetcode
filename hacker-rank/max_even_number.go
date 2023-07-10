package hacker_rank

import (
	"strconv"
)

func solution(s string) string {
	rs := ""

	arr := []rune(s)
	for i := len(arr) - 1; i >= 0; i-- {
		v, _ := strconv.Atoi(string(arr[i]))
		if v%2 == 0 {
			ui := arr[:i+1]
			rs += string(ui)
		}
	}

	return rs
}
