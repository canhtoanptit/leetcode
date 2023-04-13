package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConversion_1(t *testing.T) {
	expected := "PAHNAPLSIIGYIR"
	assert.Equal(t, expected, convert("PAYPALISHIRING", 3))
}

func TestConversion_2(t *testing.T) {
	expected := "PINALSIGYAHRPI"
	assert.Equal(t, expected, convert("PAYPALISHIRING", 4))
}
