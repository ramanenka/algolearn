package fasterallpairshortestpath

func fasterAllPairShortestPath(W [][]int) [][]int {
	L := copyMatrix(W)
	m := 1
	for m < len(W)-1 {
		L = extendPath(L, L)
		m = 2 * m
	}

	return L
}

func extendPath(L, W [][]int) [][]int {
	n := len(L)
	Ln := copyMatrix(L)
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
