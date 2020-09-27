package matrix

import (
	"testing"
)

func TestGetMaxSubmatrix(t *testing.T) {
	var tests = []struct {
		matrix [][]int32
		want int32
	}{
		{
			[][]int32{
				{1, 1, 1},
				{1, 1, 0},
				{0, 1, 0},
			},
			2,
		},
		{
			[][]int32{
				{0, 1, 1, 1},
				{1, 1, 1, 0},
				{0, 1, 0, 1},
				{0, 1, 0, 1},
			},
			2,
		},
		{
			[][]int32{
				{0, 1, 1, 1},
				{1, 1, 1, 1},
				{0, 1, 1, 1},
				{0, 1, 1, 1},
			},
			3,
		},
		{
			[][]int32{
				{0, 1, 0},
				{1, 1, 1},
				{0, 1, 0},
			},
			1,
		},
	}

	for _, tt := range tests {
		res := getMaxSubmatrix(tt.matrix)
		if res != tt.want {
			t.Errorf("got %d, want %d", res, tt.want)
		}
	}
}
