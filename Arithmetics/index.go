package main

import "fmt"

func main() {
	var a = 10
	var b = 5
	var c = a + b

	// Augmented Assignment
	var i = 10
	i += 20

	// Unary Operator
	// ++ -> a = a+1
	// -- -> a = a-1

	fmt.Println("Nilai C ", c)
	fmt.Println("Nilai A / B ", a/b)
	fmt.Println("Nilai A % B ", a%b)
	fmt.Println("Nilai C - A ", c-a)
	fmt.Println("Nilai I ", i)
	fmt.Println("Nilai I - C ", i-c)
}