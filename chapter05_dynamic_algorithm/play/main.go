package main

import (
	chapter05dynamicalgorithm "chapter05_dynamic_algorithm"
	"fmt"
)

func main() {
	d0 := 10
	d1 := 20
	d2 := 5
	d3 := 15
	d4 := 30
	M := make([][]int, 0)

	M = chapter05dynamicalgorithm.MatrixChain(M, d0, d1, d2, d3, d4)
	fmt.Println(M[0])
	fmt.Println(M[1])
	fmt.Println(M[2])
	fmt.Println(M[3])

}
