package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumOfClosedIslandsTest(t *testing.T) {
	grid := [][]int{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}}
	rs := closedIsland(grid)
	assert.Equal(t, 2, rs)
}
