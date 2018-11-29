package floydwarshall

import "math"

func floydWarshall(W [][]int) ([][]int, [][]int) {
	P := newMatrix(len(W), -1)
	for i := 0; i < len(W); i++ {
		for j := 0; j < len(W); j++ {
			if i == j || W[i][j] == math.MaxInt32 {
				continue
			}
			P[i][j] = i
		}
	}
	D := copyMatrix(W)
	for k := 0; k < len(W); k++ {
		for i := 0; i < len(W); i++ {
			for j := 0; j < len(W); j++ {
				if D[i][k]+D[k][j] < D[i][j] {
					D[i][j] = D[i][k] + D[k][j]
					P[i][j] = P[k][j]
				}
			}
		}
	}

	return D, P
}

func copyMatrix(M [][]int) [][]int {
	result := make([][]int, len(M))
	for i := range result {
		result[i] = make([]int, len(M))
	}

	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M); j++ {
			result[i][j] = M[i][j]
		}
	}
	return result
}

func newMatrix(n int, init int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = init
		}
	}
	return result
}
