package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name  string
		input args
		want  int
	}{
		{"when input is 0, should return 0", args{0}, 0},
		{"when input is 1, should return 1", args{1}, 1},
		{"when input is 2, should return 1", args{2}, 1},
		{"when input is 3, should return 2", args{3}, 2},
		{"when input is 4, should return 2", args{4}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.input.n); got != tt.want {
				assert.Equal(t, got, tt.want, "they should be equal")
			}
		})
	}
}
