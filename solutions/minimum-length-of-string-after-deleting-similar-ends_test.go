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
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minimumLength(tt.args.s), "minimumLength(%v)", tt.args.s)
		})
	}
}
