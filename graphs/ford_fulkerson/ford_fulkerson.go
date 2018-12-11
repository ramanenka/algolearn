package fordfulkerson

func fordFulkerson(C [][]int, s, t int) [][]int {
	F := make([][]int, len(C))
	for i := range F {
		F[i] = make([]int, len(C))
	}

	for {
		P := make([]int, len(C))
		for i := range P {
			P[i] = -1
		}
		V := make([]bool, len(C))
		V[s] = true
		q := []int{s}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]

			for v := 0; v < len(C); v++ {
				if V[v] {
					continue
				}

				if C[u][v] > 0 && C[u][v]-F[u][v] > 0 || C[v][u] > 0 && F[v][u] > 0 {
					P[v] = u
					V[v] = true
					q = append(q, v)
				}
			}
		}

		if P[t] == -1 {
			break
		}

		v := t
		u := P[t]
		cfp := -1
		for u != -1 {
			cfuv := 0
			if C[u][v] > 0 {
				cfuv = C[u][v] - F[u][v]
			}
			if C[v][u] > 0 {
				cfuv = F[v][u]
			}
			if cfp == -1 || cfp > cfuv {
				cfp = cfuv
			}

			v = u
			u = P[u]
		}

		v = t
		u = P[t]
		for u != -1 {
			if C[u][v] > 0 {
				F[u][v] += cfp
			}
			if C[v][u] > 0 {
				F[v][u] -= cfp
			}
			v = u
			u = P[u]
		}
	}

	return F
}
