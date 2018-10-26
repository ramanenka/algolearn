package graphs

import (
	"math"
)

/*
Bellman Ford Algorithm
https://en.wikipedia.org/wiki/Bellman%E2%80%93Ford_algorithm
*/

type bellmanFordVertex struct {
	name      string
	d         int
	dExpected int
	p         *bellmanFordVertex
	adj       []*bellmanFordEdge
}

type bellmanFordEdge struct {
	w int
	v *bellmanFordVertex
}

func bellmanFord(vertices []*bellmanFordVertex) bool {
	for _, v := range vertices {
		v.d = math.MaxInt32
	}
	vertices[0].d = 0

	for i := 1; i <= len(vertices)-1; i++ {
		for _, u := range vertices {
			for _, e := range u.adj {
				v := e.v
				if u.d+e.w < v.d {
					v.d = u.d + e.w
					v.p = u
				}
			}
		}
	}

	for _, u := range vertices {
		for _, e := range u.adj {
			v := e.v
			if v.d > u.d+e.w {
				return false
			}
		}
	}

	return true
}
