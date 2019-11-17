package main

import (
	"fmt"
	"strconv"
)

func sum(slice []int) int {
	total := 0
	for i := 0; i < len(slice); i++ {
		total = total + slice[i]
	}
	return total
}

func solution(N int) []int {
	array := make([]int, N)
	total := 0
	for i := 0; i < N-1; i++ {
		array[i] = i
		total = total + i
	}
	array[N-1] = -(total)
	return array
}

func main() {
	fmt.Println(solution(4))
	fmt.Println("Sum: " + strconv.Itoa(sum(solution(4))))
	fmt.Println(solution(3))
	fmt.Println("Sum: " + strconv.Itoa(sum(solution(3))))
}
