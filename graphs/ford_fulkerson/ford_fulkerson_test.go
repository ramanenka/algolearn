package fordfulkerson

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/ramanenka/algolearn/graphs"
)

func Test_fordFulkerson(t *testing.T) {
	type args struct {
		C [][]int
		s int
		t int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case1",
			args: args{
				C: [][]int{
					[]int{0, 16, 13, 0, 0, 0},
					[]int{0, 0, 0, 12, 0, 0},
					[]int{0, 4, 0, 0, 14, 0},
					[]int{0, 0, 9, 0, 0, 20},
					[]int{0, 0, 0, 7, 0, 4},
					[]int{0, 0, 0, 0, 0, 0},
				},
				s: 0,
				t: 5,
			},
			want: [][]int{
				[]int{0, 12, 11, 0, 0, 0},
				[]int{0, 0, 0, 12, 0, 0},
				[]int{0, 0, 0, 0, 11, 0},
				[]int{0, 0, 0, 0, 0, 19},
				[]int{0, 0, 0, 7, 0, 4},
				[]int{0, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fordFulkerson(tt.args.C, tt.args.s, tt.args.t)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fordFulkerson() = %v, want %v", got, tt.want)
			}

			sb := &strings.Builder{}
			sb.WriteString("digraph G {\n")
			sb.WriteString(fmt.Sprintf("\tlabel=\"Maximum Flow found by Ford-Fulkerson method\";\n"))
			for u := range tt.args.C {
				for v := range tt.args.C {
					if tt.args.C[u][v] == 0 {
						continue
					}

					sb.WriteString(fmt.Sprintf("\t%d -> %d ", u, v))
					sb.WriteString("[")
					sb.WriteString(fmt.Sprintf("label=\"%d/%d\"", got[u][v], tt.args.C[u][v]))
					if got[u][v] > 0 {
						sb.WriteString(fmt.Sprintf(",penwidth=%d", got[u][v]>>1))
					} else {
						sb.WriteString(",style=dotted")
					}

					sb.WriteString("];\n")
				}
			}

			sb.WriteString("}\n")
			graphs.Dotviz(tt.name, sb.String())
		})
	}
}
