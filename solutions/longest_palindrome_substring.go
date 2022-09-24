package leetcode

func longestPalindromeBrutalForce(s string) string {
	// find all possible substring
	// for from start, end of substring. if all match ->Palindrome
	allSubString := findAllSubstring(s)
	rs := ""
	lengtS := 0
	for _, e := range allSubString {
		if isPalindromeString(e) {
			if lengtS < len(e) {
				lengtS = len(e)
				rs = e
			}
		}
	}
	return rs
}

func findAllSubstring(s string) []string {
	n := len(s)
	subS := make([]string, 0)
	for i := 0; i < n; i++ {
		subS = append(subS, s[:i+1])
		ss := s[i+1:]
		var sj []string
		for j := 0; j < len(ss); j++ {
			sj = append(sj, ss[:j+1])
		}
		subS = append(subS, sj...)
	}

	return subS
}

func isPalindromeString(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}
