package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}

	fmt.Println(reduceMultiply(arr))
}

func reduceMultiply(arr []int) int {
	result := 1
	for _, v := range arr {
		result *= v
	}

	return result
}
