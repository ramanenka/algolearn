package dijkstras

import (
	"math"
)

type vertex struct {
	name      string
	d         int
	dExpected int
	p         *vertex
	adj       []*edge
	heapIDx   int
}

type edge struct {
	v *vertex
	w int
}

func dijkstra(vertices []*vertex) {
	for _, v := range vertices {
		v.d = math.MaxInt32
	}
	vertices[0].d = 0

	h := []*vertex{}
	for i, v := range vertices {
		v.heapIDx = i
		h = append(h, v)
	}
	buildHeap(h)

	for len(h) > 0 {
		var u *vertex
		h, u = heapExtract(h)
		for _, e := range u.adj {
			v := e.v
			if u.d+e.w < v.d {
				v.d = u.d + e.w
				v.p = u
				heapDecrease(h, v.heapIDx)
			}
		}
	}
}

func heapify(h []*vertex, i int) {
	l := left(i)
	r := right(i)
	min := i

	if l < len(h) && h[l].d < h[min].d {
		min = l
	}

	if r < len(h) && h[r].d < h[min].d {
		min = r
	}

	if i != min {
		h[i], h[min] = h[min], h[i]
		h[min].heapIDx, h[i].heapIDx = h[i].heapIDx, h[min].heapIDx
		heapify(h, min)
	}
}

func buildHeap(h []*vertex) {
	for i := len(h) / 2; i >= 0; i-- {
		heapify(h, i)
	}
}

func heapExtract(h []*vertex) ([]*vertex, *vertex) {
	v := h[0]
	h[0] = h[len(h)-1]
	h[0].heapIDx = 0
	h = h[:len(h)-1]
	heapify(h, 0)
	return h, v
}

func heapDecrease(h []*vertex, idx int) {
	p := parent(idx)
	if p < 0 {
		return
	}

	if h[idx].d < h[p].d {
		h[idx], h[p] = h[p], h[idx]
		h[idx].heapIDx, h[p].heapIDx = h[p].heapIDx, h[idx].heapIDx
	}
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

func parent(i int) int {
	return (i - 1) / 2
}
