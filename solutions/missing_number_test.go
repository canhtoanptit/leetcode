package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_missingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "correct",
			args: args{
				nums: []int{3, 0, 1},
			},
			want: 2,
		},
		{
			name: "correct1",
			args: args{
				nums: []int{0, 1},
			},
			want: 2,
		},
		{
			name: "correct2",
			args: args{
				nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			},
			want: 8,
		},
		{
			name: "correct3",
			args: args{
				nums: []int{6, 4, 2, 3, 5, 7, 8, 1},
			},
			want: 0,
		},
		{
			name: "correct4",
			args: args{
				nums: []int{0, 2, 3},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, missingNumber(tt.args.nums), "missingNumber(%v)", tt.args.nums)
		})
	}
}
