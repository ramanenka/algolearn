package floydwarshall

func floydWarshall(W [][]int) [][]int {
	D := copyMatrix(W)
	for k := 0; k < len(W); k++ {
		for i := 0; i < len(W); i++ {
			for j := 0; j < len(W); j++ {
				if D[i][k]+D[k][j] < D[i][j] {
					D[i][j] = D[i][k] + D[k][j]
				}
			}
		}
	}

	return D
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
