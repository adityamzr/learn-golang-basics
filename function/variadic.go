package main

import "fmt"

func sumAll(numbers ...int) int  {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	sum := sumAll(10, 20, 30, 40)
	fmt.Println(sum)
	// atau
	numbers := []int{10, 10, 10}
	fmt.Println(sumAll(numbers...))
}