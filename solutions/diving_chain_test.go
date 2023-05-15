package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolutionDivingChain(t *testing.T) {
	input := []int{5, 2, 4, 6, 3, 7}
	rs := SolutionDivingChain(input)
	assert.Equal(t, 5, rs)
}
