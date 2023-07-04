package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleNum(t *testing.T) {
	i := []int{2, 2, 3, 2}
	o := singleNumber(i)
	assert.Equal(t, 3, o)
}

func TestSingleNum1(t *testing.T) {
	i := []int{0, 1, 0, 1, 0, 1, 99}
	o := singleNumber(i)
	assert.Equal(t, 99, o)
}

func TestSingleNum2(t *testing.T) {
	i := []int{30000, 500, 100, 30000, 100, 30000, 100}
	o := singleNumber(i)
	assert.Equal(t, 500, o)
}
