package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minimumLength(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TC1",
			args: args{
				s: "ca",
			},
			want: 2,
		},
		{
			name: "TC2",
			args: args{
				s: "cabaabac",
			},
			want: 0,
		},
		{
			name: "TC3",
			args: args{
				s: "aabccabba",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minimumLength(tt.args.s), "minimumLength(%v)", tt.args.s)
		})
	}
}
