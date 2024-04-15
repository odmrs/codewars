package main

import "fmt"

func main() {
	fmt.Println(SquareSum([]int{1, 2, 2}))
}

func SquareSum(numbers []int) (res int) {
	for _, num := range numbers {
		res += num * num
	}
	return
}
