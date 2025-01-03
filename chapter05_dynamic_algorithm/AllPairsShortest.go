package chapter05dynamicalgorithm

import (
	"fmt"
	"math"
)

const INF = math.MaxInt / 2

// math.MaxInt을 그대로 사용할 경우 오버플로우 날 확률이 있음.

var Map = [][]int{
	{0, 4, 2, 5, INF},
	{INF, 0, 1, INF, 4},
	{1, 3, 0, 1, 2},
	{-2, INF, INF, 0, 2},
	{INF, -3, 3, 1, 0},
}

func AllPairsShortest(Map [][]int) {

	col := len(Map)

	min := func(New1 int, New2 int, original *int) {

		if *original > (New1 + New2) {
			*original = New1 + New2
		}
	}

	for i := 0; i < col; i++ {
		for ii := 0; ii < col; ii++ {
			if i == ii {
				continue
			}
			for iii := 0; iii < col; iii++ {

				if iii == i || iii == ii {
					continue
				}

				original := Map[ii][iii]
				fmt.Printf("min{Map[%d,%d],Map[%d,%d]+Map[%d,%d]}", ii, iii, ii, i, i, iii)
				min(Map[ii][i], Map[i][iii], &original)
				fmt.Printf("=Map{%d,%d+%d}=%d\n", Map[ii][iii], Map[ii][i], Map[i][iii], original)
				Map[ii][iii] = original
			}
			fmt.Println(ii, Map[ii])
		}
		fmt.Println()

	}
}
