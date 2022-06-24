package main

import "fmt"

func main() {
	res := evenOrOdd(10)
	fmt.Println("The number is: ", res)
}

func evenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}
