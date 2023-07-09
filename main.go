package main

import "fmt"

func sum(n ...int) int {
	total := 0

	for _, v := range n {
		total += v
	}
	return total
}

func subtract(n ...int) int {
	total := n[0]

	for _, v := range n[1:] {
		total = total - v
	}
	return total
}

func multiply(n ...int) int {
	total := 1

	for _, v := range n {
		total *= v
	}
	return total
}

func divide(n ...float64) float64 {
	total := n[0]

	for _, v := range n[1:] {
		total = total / v
	}
	return total
}

func main() {
	fmt.Println(sum(43, 2, 13))
	fmt.Println(subtract(21, 7, 2))
	fmt.Println(multiply(10, 2, 3, 2))
	fmt.Println(divide(100, 15, 2))
}