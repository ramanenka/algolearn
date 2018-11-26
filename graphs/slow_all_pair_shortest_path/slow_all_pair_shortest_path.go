package slowallpairshortestpath

import (
	"fmt"
)

func slowAllPairShortestPath(W [][]int) ([][]int, [][]int) {
	L := copyMatrix(W)
	P := newMatrix(len(W), -1)
	for i, row := range P {
		for j := 0; j < len(row); j++ {
			P[i][j] = i
		}
	}
	for i := 2; i <= len(W)-1; i++ {
		L = extendPath(L, W, P)
	}

	return L, P
}

func extendPath(L, W, P [][]int) [][]int {
	n := len(L)
	Ln := copyMatrix(L)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if L[i][k]+W[k][j] < Ln[i][j] {
					Ln[i][j] = L[i][k] + W[k][j]
					P[i][j] = k
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

func printMatrix(M [][]int) {
	for _, row := range M {
		for _, c := range row {
			fmt.Printf("%3d ", c)
		}
		fmt.Println()
	}
}
