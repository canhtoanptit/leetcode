package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	rs := lengthOfLongestSubstringWithSlidingWindow(s)
	assert.Equal(t, 3, rs)
}

func TestLengthOfLongestSubstring1(t *testing.T) {
	s := "bbbbb"
	rs := lengthOfLongestSubstringWithSlidingWindow(s)
	assert.Equal(t, 1, rs)
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	s := "pwwkew"
	rs := lengthOfLongestSubstringWithSlidingWindow(s)
	assert.Equal(t, 3, rs)
}

func TestLengthOfLongestSubstring3(t *testing.T) {
	s := "ckilbkd"
	rs := lengthOfLongestSubstringWithSlidingWindow(s)
	assert.Equal(t, 5, rs)
}

func TestLengthOfLongestSubstring4(t *testing.T) {
	s := "dvdf"
	rs := lengthOfLongestSubstringWithSlidingWindow(s)
	assert.Equal(t, 3, rs)
}
