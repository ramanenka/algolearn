package floydwarshall

func transitiveClosure(D [][]bool) [][]bool {
	for k := 0; k < len(D); k++ {
		for i := 0; i < len(D); i++ {
			for j := 0; j < len(D); j++ {
				D[i][j] = D[i][j] || D[i][k] && D[k][j]
			}
		}
	}
	return D
}
