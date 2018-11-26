package fasterallpairshortestpath

import (
	"math"
	"reflect"
	"testing"
)

func Test_fasterAllPairShortestPath(t *testing.T) {
	inf := math.MaxInt32
	type args struct {
		W [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{
				W: [][]int{
					[]int{0, 3, 8, inf, -4},
					[]int{inf, 0, inf, 1, 7},
					[]int{inf, 4, 0, inf, inf},
					[]int{2, inf, -5, 0, inf},
					[]int{inf, inf, inf, 6, 0},
				},
			},
			want: [][]int{
				[]int{0, 1, -3, 2, -4},
				[]int{3, 0, -4, 1, -1},
				[]int{7, 4, 0, 5, 3},
				[]int{2, -1, -5, 0, -2},
				[]int{8, 5, 1, 6, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			W := tt.args.W
			L := fasterAllPairShortestPath(W)
			if !reflect.DeepEqual(L, tt.want) {
				t.Errorf("slowAllPairShortestPath() got = %v, want %v", L, tt.want)
			}
		})
	}
}
