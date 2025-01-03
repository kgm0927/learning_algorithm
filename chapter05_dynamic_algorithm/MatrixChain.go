package chapter05dynamicalgorithm

import (
	"math"
)

// MatrixChain 함수는 행렬 곱셈의 최소 비용을 계산합니다.
func MatrixChain(M [][]int, d ...int) [][]int {

	d_len := len(d) - 1

	// 만약 M이 제공되지 않으면 M을 초기화합니다.
	if len(M) == 0 {
		M = make([][]int, d_len)
		for i := 0; i < d_len; i++ {
			M[i] = make([]int, d_len)
		}
	}

	// matrix_multiple 함수는 행렬 차원 d1 x d2 x d3을 곱하는 비용을 계산합니다.
	matrix_multiple := func(d1, d2, d3 int) int {
		return d1 * d2 * d3
	}

	// M 행렬의 대각선은 0으로 초기화합니다. (한 개의 행렬은 곱셈이 필요없음)
	for i := 0; i < d_len; i++ {
		M[i][i] = 0
	}

	for length := 1; length < len(d); length++ {
		for i := 0; i < d_len-length; i++ {
			j := i + length
			M[i][j] = math.MaxInt32

			for k := i; k < j; k++ { // 연속된 행렬의 곱셈을 횟수만큼 실행한 후 제일 작은 값을 보냄
				q := M[i][k] + M[k+1][j] + matrix_multiple(d[i], d[k+1], d[j+1])

				if q < M[i][j] {
					M[i][j] = q
				}
			}
		}
	}

	return M
}
