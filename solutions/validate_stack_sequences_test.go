package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateStackSequences1(t *testing.T) {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	assert.True(t, validateStackSequences(pushed, popped))
}

func TestValidateStackSequences2(t *testing.T) {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 3, 5, 1, 2}
	assert.False(t, validateStackSequences(pushed, popped))
}
