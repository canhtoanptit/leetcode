package leetcode

import (
	"fmt"
	"math/big"
)

func maxValue(events [][]int, k int) int {
	mapOfDays := make(map[int][][]int)
	mMaxValue := make(map[int]int)
	var limit int
	for _, e := range events {
		startDay := e[0]
		endDay := e[1]
		if endDay >= limit {
			limit = endDay
		}
		dd, ok := mapOfDays[startDay]
		if !ok {
			mapOfDays[startDay] = [][]int{e}
		} else {
			dd = append(dd, e)
			mapOfDays[startDay] = dd
		}
	}

	for i := 1; i <= limit; i++ {
		eventsByDay := mapOfDays[i]
		for _, event := range eventsByDay {
			mMaxValue[i] = event[0]
		}
	}

	return 1
}

func getMaxOf(start, k int, mMaxValue map[int]int) int {
	return 1
}

func plusMinus(arr []int) {
	var p, n, z int
	total := len(arr)
	for _, e := range arr {
		if e > 0 {
			p++
		} else if e == 0 {
			z++
		} else {
			n++
		}
	}

	fmt.Println(fmt.Sprintf("%.6f", float64(p)/float64(total)))
}

func staircase(n int) {
	// Write your code here
	for i := 1; i <= n; i++ {
		fmt.Println(print(n-i, i))
	}
}

func print(a, b int) string {
	rs := ""
	for a > 0 {
		rs += " "
		a--
	}

	for b > 0 {
		rs += "#"
		b--
	}

	return rs
}

func extraLongFactorials(n int) {
	// Write your code here
	var fact big.Int
	fact.MulRange(1, int64(n))
	fmt.Println(fact)
}
