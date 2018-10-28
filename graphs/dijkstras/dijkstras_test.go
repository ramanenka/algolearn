package dijkstras

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ramanenka/algolearn/graphs"
)

func Test_dijkstra(t *testing.T) {
	type args struct {
		vertices []*vertex
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				vertices: func() []*vertex {
					s := &vertex{name: "s", dExpected: 0}
					t := &vertex{name: "t", dExpected: 8}
					y := &vertex{name: "y", dExpected: 5}
					x := &vertex{name: "x", dExpected: 9}
					z := &vertex{name: "z", dExpected: 7}
					s.adj = []*edge{
						&edge{v: t, w: 10},
						&edge{v: y, w: 5},
					}
					t.adj = []*edge{
						&edge{v: x, w: 1},
						&edge{v: y, w: 2},
					}
					y.adj = []*edge{
						&edge{v: t, w: 3},
						&edge{v: x, w: 9},
						&edge{v: z, w: 2},
					}
					x.adj = []*edge{
						&edge{v: z, w: 4},
					}
					z.adj = []*edge{
						&edge{v: s, w: 7},
						&edge{v: x, w: 6},
					}

					return []*vertex{s, t, y, x, z}
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dijkstra(tt.args.vertices)
			for _, v := range tt.args.vertices {
				if v.d != v.dExpected {
					t.Errorf("expected vertex %s d to be %d, got %d", v.name, v.dExpected, v.d)
				}
			}

			sb := &strings.Builder{}
			sb.WriteString("digraph G {\n")
			for _, u := range tt.args.vertices {
				sb.WriteString(fmt.Sprintf("\t%s [label=\"%s/%d\"];\n", u.name, u.name, u.d))
				for _, e := range u.adj {
					v := e.v
					sb.WriteString(fmt.Sprintf("\t%s -> %s ", u.name, v.name))
					sb.WriteString("[")
					sb.WriteString(fmt.Sprintf("label=\"%d\"", e.w))
					if v.p == u {
						sb.WriteString(",style=bold")
					}
					sb.WriteString("];\n")
				}
			}
			sb.WriteString("}")
			graphs.Dotviz("dijkstras_"+tt.name, sb.String())
		})
	}
}
