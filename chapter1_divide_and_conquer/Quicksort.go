package main

import (
	"fmt"
	"math/rand/v2"
)

func arrange_the_slice(start, end int, slice []int) {
	if start >= end {
		return
	}

	// 피벗을 중앙 값으로 설정
	pivotIndex := (start + end) / 2
	pivot := slice[pivotIndex]

	// 피벗을 첫 번째 값과 교환
	slice[pivotIndex], slice[start] = slice[start], slice[pivotIndex]

	// 배열을 분할
	i := start + 1
	ii_n := end
	for i <= ii_n {
		if slice[i] <= pivot {
			i++
		} else if slice[ii_n] > pivot {
			ii_n--
		} else {
			slice[i], slice[ii_n] = slice[ii_n], slice[i]
			i++
			ii_n--
		}
	}

	// 피벗을 분할된 중앙에 위치
	slice[start], slice[ii_n] = slice[ii_n], slice[start]

	// 두 부분을 각각 재귀적으로 정렬
	arrange_the_slice(start, ii_n-1, slice)
	arrange_the_slice(ii_n+1, end, slice)
}

func main() {
	slice := []int{}

	for i := 0; i < 100; i++ {
		slice = append(slice, rand.IntN(100))
	}
	fmt.Println(slice)

	arrange_the_slice(0, 99, slice)

	fmt.Println(slice)
}
