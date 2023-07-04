package leetcode

func singleNumber(nums []int) int {
	mm := make(map[int]int)

	for _, e := range nums {
		v, ok := mm[e]
		if !ok {
			mm[e] = 1
		} else {
			v += 1
			mm[e] = v
		}
	}

	for k, v := range mm {
		if v == 1 {
			return k
		}
	}

	return -1
}
