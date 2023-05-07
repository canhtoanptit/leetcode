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
