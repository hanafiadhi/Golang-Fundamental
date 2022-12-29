package main

import "fmt"

func sumAll(number ...int) int {
	total := 0

	for _, value := range number {
		total += value
	}
	return total
}

func main() {
	fmt.Println(sumAll(10, 10, 101, 10, 10, 3123, 1023, 1312, 123, 34, 5, 7))

	slice := []int{1, 2, 3, 4, 5, 6}

	total := sumAll(slice...)

	fmt.Println(total)
}
