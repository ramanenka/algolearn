package slowallpairshortestpath

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"testing"

	"github.com/ramanenka/algolearn/graphs"
)

func Test_slowAllPairShortestPath(t *testing.T) {
	inf := math.MaxInt32
	type args struct {
		W [][]int
	}
	type tovis struct {
		u int
		v int
	}
	tests := []struct {
		name  string
		args  args
		want  [][]int
		want1 [][]int
		tovis tovis
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
			want1: [][]int{
				[]int{0, 2, 3, 4, 0},
				[]int{3, 1, 3, 1, 0},
				[]int{3, 2, 2, 1, 0},
				[]int{3, 2, 3, 3, 0},
				[]int{3, 2, 3, 4, 4},
			},
			tovis: tovis{
				u: 0,
				v: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			W := tt.args.W
			L, P := slowAllPairShortestPath(W)
			if !reflect.DeepEqual(L, tt.want) {
				t.Errorf("slowAllPairShortestPath() got = %v, want %v", L, tt.want)
			}
			if !reflect.DeepEqual(P, tt.want1) {
				t.Errorf("slowAllPairShortestPath() got = %v, want %v", P, tt.want1)
			}

			sb := &strings.Builder{}
			sb.WriteString("digraph G {\n")
			sb.WriteString(fmt.Sprintf(
				"\tlabel=\"Shortest path between %d and %d is of length %d\";\n",
				tt.tovis.u,
				tt.tovis.v,
				L[tt.tovis.u][tt.tovis.v],
			))
			for u, row := range W {
				for v, w := range row {
					if u == v || w == inf {
						continue
					}

					sb.WriteString(fmt.Sprintf("\t%d -> %d ", u, v))
					sb.WriteString("[")
					sb.WriteString(fmt.Sprintf("label=\"%d\"", w))

					isbold := false
					up := tt.tovis.u
					vp := tt.tovis.v
					for vp != up {
						p := P[up][vp]
						if u == p && vp == v {
							isbold = true
							break
						}
						vp = p
					}

					if isbold {
						sb.WriteString(",style=bold")
					}

					sb.WriteString("];\n")
				}
			}

			sb.WriteString("}\n")
			graphs.Dotviz(tt.name, sb.String())
		})
	}
}
