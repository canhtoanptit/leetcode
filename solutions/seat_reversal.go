package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func SolutionSeat(N int, S string) int {
	// Implement your solution here
	reservedSeats := toReservationIndex(S)
	fmt.Printf("%+v", reservedSeats)
	rows := map[int]int{}

	m1 := 0b111100
	m2 := 0b11110000
	m3 := 0b1111000000

	for _, seat := range reservedSeats {
		i := seat[0]
		j := seat[1]

		rows[i] |= 1 << j
	}

	count := 0

	for _, row := range rows {
		if row&m1 == 0 {
			count++
			if row&m3 == 0 {
				count++
			}
		} else if row&m2 == 0 {
			count++
		} else if row&m3 == 0 {
			count++
		}
	}

	count += (N - len(rows)) * 2
	return count
}

func toReservationIndex(S string) [][]int {
	mapPosition := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7,
		"H": 8,
		"J": 9,
		"K": 10,
	}
	var rs [][]int
	if "" == S {
		return rs
	}
	positions := strings.Split(S, " ")
	for _, v := range positions {
		size := len(v)
		s := v[:size-1]
		iv, _ := strconv.Atoi(s)
		rs = append(rs, []int{iv, mapPosition[v[size-1:]]})
	}

	return rs
}
