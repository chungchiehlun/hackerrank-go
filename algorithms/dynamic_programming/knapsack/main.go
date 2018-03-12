package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var testCases int
	fmt.Scan(&testCases)

	for t := 0; t < testCases; t++ {
		var N, K int
		fmt.Scan(&N)
		fmt.Scan(&K)

		K = K + 1

		A := make([]int, N)
		for i := range A {
			fmt.Scan(&A[i])
		}

		T := make([][]int, N)
		for i := 0; i < N; i++ {
			T[i] = make([]int, K)
		}

		for n := 0; n < N; n++ {
			for k := 0; k < K; k++ {
				if A[n] > k {
					if n == 0 {
						T[n][k] = 0
					} else {
						T[n][k] = T[n-1][k]
					}
				} else {
					if n == 0 {
						T[n][k] = A[n] + T[n][k-A[n]]
					} else {
						T[n][k] = max(A[n]+T[n][k-A[n]], T[n-1][k])
					}
				}
			}
		}

		fmt.Println(T[N-1][K-1])
	}
}
