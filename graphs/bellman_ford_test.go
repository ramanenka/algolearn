package graphs

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func Test_bellmanFord(t *testing.T) {
	type args struct {
		vertices []*bellmanFordVertex
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				vertices: func() []*bellmanFordVertex {
					s := &bellmanFordVertex{name: "s", dExpected: 0}
					t := &bellmanFordVertex{name: "t", dExpected: 2}
					y := &bellmanFordVertex{name: "y", dExpected: 7}
					x := &bellmanFordVertex{name: "x", dExpected: 4}
					z := &bellmanFordVertex{name: "z", dExpected: -2}
					s.adj = []*bellmanFordEdge{
						&bellmanFordEdge{v: t, w: 6},
						&bellmanFordEdge{v: y, w: 7},
					}
					t.adj = []*bellmanFordEdge{
						&bellmanFordEdge{v: x, w: 5},
						&bellmanFordEdge{v: z, w: -4},
						&bellmanFordEdge{v: y, w: 8},
					}
					x.adj = []*bellmanFordEdge{
						&bellmanFordEdge{v: t, w: -2},
					}
					y.adj = []*bellmanFordEdge{
						&bellmanFordEdge{v: x, w: -3},
						&bellmanFordEdge{v: z, w: 9},
					}
					z.adj = []*bellmanFordEdge{
						&bellmanFordEdge{v: s, w: 2},
						&bellmanFordEdge{v: x, w: 7},
					}

					return []*bellmanFordVertex{s, t, y, x, z}
				}(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bellmanFord(tt.args.vertices); got != tt.want {
				t.Errorf("bellmanFord() = %v, want %v", got, tt.want)
			}
			for _, v := range tt.args.vertices {
				if v.d != v.dExpected {
					t.Errorf("expected d of vertex %s is %d, but got %d", v.name, v.dExpected, v.d)
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
			dotviz("bellman_ford_"+tt.name, sb.String())
		})
	}
}

func dotviz(name, data string) {
	cmd := exec.Command("dot", "-Tpng")

	f, err := os.Create(fmt.Sprintf("%s.png", name))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	cmd.Stdin = strings.NewReader(data)
	cmd.Stdout = f
	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
