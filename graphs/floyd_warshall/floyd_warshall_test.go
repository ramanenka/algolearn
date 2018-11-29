package floydwarshall

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"testing"

	"github.com/ramanenka/algolearn/graphs"
)

func Test_floydWarshall(t *testing.T) {
	inf := math.MaxInt32
	type args struct {
		W [][]int
	}
	type tovis struct {
		u, v int
	}
	tests := []struct {
		name  string
		args  args
		want  [][]int
		P     [][]int
		tovis []tovis
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
			P: [][]int{
				[]int{-1, 3, 0, 2},
				[]int{1, -1, 0, 2},
				[]int{1, 3, -1, 2},
				[]int{1, 3, 0, -1},
			},
			tovis: []tovis{{u: 2, v: 0}, {u: 1, v: 3}},
		},
		{
			name: "case2",
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
			P: [][]int{
				[]int{-1, 2, 3, 4, 0},
				[]int{3, -1, 3, 1, 0},
				[]int{3, 2, -1, 1, 0},
				[]int{3, 2, 3, -1, 0},
				[]int{3, 2, 3, 4, -1},
			},
			tovis: []tovis{{u: 0, v: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, P := floydWarshall(tt.args.W)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floydWarshall() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(P, tt.P) {
				t.Errorf("floydWarshall() = _, %v, want _, %v", P, tt.P)
			}

			for _, tovis := range tt.tovis {
				sb := &strings.Builder{}
				sb.WriteString("digraph G {\n")
				sb.WriteString(fmt.Sprintf(
					"\tlabel=\"Shortest path between %d and %d is of length %d\";\n",
					tovis.u,
					tovis.v,
					got[tovis.u][tovis.v],
				))
				for u, row := range tt.args.W {
					for v, w := range row {
						if u == v || w == inf {
							continue
						}

						sb.WriteString(fmt.Sprintf("\t%d -> %d ", u, v))
						sb.WriteString("[")
						sb.WriteString(fmt.Sprintf("label=\"%d\"", w))

						isbold := false
						up := tovis.u
						vp := tovis.v
						for vp != up {
							p := P[up][vp]
							if u == p && vp == v {
								isbold = true
								break
							}
							vp = p
						}

						if isbold {
							sb.WriteString(",style=bold,color=red")
						}

						sb.WriteString("];\n")
					}
				}

				sb.WriteString("}\n")
				graphs.Dotviz(fmt.Sprintf("%s_%d_%d", tt.name, tovis.u, tovis.v), sb.String())
			}
		})
	}
}
