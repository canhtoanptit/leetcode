package hacker_rank

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClimbingLeaderboard(t *testing.T) {
	o := climbingLeaderboard([]int32{100, 90, 90, 80}, []int32{70, 80, 105})
	assert.Equal(t, []int32{4, 3, 1}, o)
}
