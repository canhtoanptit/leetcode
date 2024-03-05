package leetcode

func minimumLength(s string) int {
	left := 0
	right := len(s) - 1
	for left < right && s[left] == s[right] {
		char := s[left]
		for left <= right && s[left] == char {
			left++
		}

		for right >= left && s[right] == char {
			right--
		}
	}

	return right - left + 1
}
