package main

import "fmt"

func result(a int, b int) int {
	return a + b
}

func main()  {
	var hasil int = 5

	add := result(1, 2)

	hasil += add

	fmt.Println("Hasilnya adalah", hasil)
}