package leetcode

import (
	"fmt"
	"reflect"
	"sort"
)

func threeSum(nums []int) [][]int {
	var rs [][]int
	// naive way
	// find all possible triples
	// remove duplicate triples
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					e := []int{nums[i], nums[j], nums[k]}
					less := func(a, b int) bool { return a < b }
					sort.Slice(e, less)
					fmt.Printf("%+v", e)
					if len(rs) == 0 || !checkDuplicateElement(rs, e) {
						rs = append(rs, e)
					}
				}
			}
		}
	}
	return rs
}

func checkDuplicateElement(source [][]int, check []int) bool {
	for _, v := range source {
		if containAllElement(v, check) {
			return true
		}
	}
	return false
}

func containAllElement(source, check []int) bool {
	return reflect.DeepEqual(source, check)
}
