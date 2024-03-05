package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_bagOfTokensScore(t *testing.T) {
	type args struct {
		tokens []int
		power  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Testcase 1",
			args: args{
				tokens: []int{100},
				power:  50,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, bagOfTokensScore(tt.args.tokens, tt.args.power), "bagOfTokensScore(%v, %v)", tt.args.tokens, tt.args.power)
		})
	}
}
