package main

import "fmt"

func main() {
	var name1 = "Nay"
	var name2 = "Nava"

	var result = name1 < name2
	fmt.Println(result)

	var val1 = 100
	var val2 = 200

	fmt.Println(val1 > val2)
	fmt.Println(val1 < val2)
	fmt.Println(val1 == val2)
	fmt.Println(val1 != val2)
}