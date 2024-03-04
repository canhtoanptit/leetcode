package gccr

import "testing"

func TestF(t *testing.T) {
	F()
}

func TestF2(t *testing.T) {
	DoWork()
}

func TestNewRunStream(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewRunStream()
		})
	}
}

func TestCheckStatus(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoWork()
		})
	}
}
