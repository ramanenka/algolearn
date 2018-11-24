package slowallpairshortestpath

import (
	"math"
)

func slowAllPairShortestPath(W [][]int) [][]int {
	L := copyMatrix(W)
	for i := 2; i <= len(W)-1; i++ {
		L = extendPath(L, W)
	}
	return L
}

func extendPath(L, W [][]int) [][]int {
	n := len(L)
	Ln := newMatrix(n, math.MaxInt32)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if L[i][k]+W[k][j] < Ln[i][j] {
					Ln[i][j] = L[i][k] + W[k][j]
				}
			}
		}
	}
	return Ln
}

func newMatrix(n, init int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
		for j := range result[i] {
			result[i][j] = init
		}
	}
	return result
}

func copyMatrix(M [][]int) [][]int {
	result := make([][]int, len(M))
	for i := range result {
		result[i] = make([]int, len(M))
		for j := range result[i] {
			result[i][j] = M[i][j]
		}
	}
	return result
}
