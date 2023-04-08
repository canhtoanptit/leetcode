package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumEnclavesTest(t *testing.T) {
	grid := [][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}
	rs := numEnclaves(grid)
	assert.Equal(t, 3, rs)
}
