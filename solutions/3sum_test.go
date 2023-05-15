package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test3sum(t *testing.T) {
	in := []int{-1, 0, 1, 2, -1, -4}
	out := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	assert.Equal(t, out, threeSum(in))
}

func Test3sum1(t *testing.T) {
	in := []int{0, 1, 1}
	var out [][]int
	assert.Equal(t, out, threeSum(in))
}

func Test3sum2(t *testing.T) {
	in := []int{0, 0, 0}
	out := [][]int{{0, 0, 0}}
	assert.Equal(t, out, threeSum(in))
}

func Test3sum3(t *testing.T) {
	in := []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}
	out := [][]int{{-5, 1, 4}, {-4, 0, 4}, {-4, 1, 3}, {-2, -2, 4}, {-2, 1, 1}, {0, 0, 0}}
	assert.Equal(t, out, threeSum(in))
}
