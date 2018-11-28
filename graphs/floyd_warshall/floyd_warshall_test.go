package floydwarshall

import (
	"math"
	"reflect"
	"testing"
)

func Test_floydWarshall(t *testing.T) {
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
					[]int{0, inf, -2, inf},
					[]int{4, 0, 3, inf},
					[]int{inf, inf, 0, 2},
					[]int{inf, -1, inf, 0},
				},
			},
			want: [][]int{
				[]int{0, -1, -2, 0},
				[]int{4, 0, 2, 4},
				[]int{5, 1, 0, 2},
				[]int{3, -1, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floydWarshall(tt.args.W); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floydWarshall() = %v, want %v", got, tt.want)
			}
		})
	}
}
