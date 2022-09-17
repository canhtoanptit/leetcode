package leetcode

import (
	"math"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	chars := []rune(s)
	if len(chars) == 0 {
		return 0
	}
	var subString string
	var rs int
	lengOfString := len(chars)
	for i := 0; i < lengOfString; i++ {
		if !strings.ContainsRune(subString, chars[i]) {
			subString += string(chars[i])
			rs++
		} else {
			newS := string(chars[1:lengOfString])
			sub := lengthOfLongestSubstring(newS)
			if rs < sub {
				return sub
			} else {
				return rs
			}
		}
	}
	return rs
}

func lengthOfLongestSubstringWithSlidingWindow(s string) int {
	n := len(s)
	res := 0
	mp := make(map[uint8]int, 0)

	for i, j := 0, 0; j < n; j++ {
		if mp[s[j]] > 0 {
			i = int(math.Max(float64(mp[s[j]]), float64(i)))
		}
		res = int(math.Max(float64(res), float64(j-i+1)))
		mp[s[j]] = j + 1
	}

	return res
}
