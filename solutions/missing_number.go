package leetcode

func missingNumber(nums []int) int {
	missing := len(nums)
	for i, e := range nums {
		missing ^= i ^ e
	}

	return missing
}
