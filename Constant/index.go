package main

import "fmt"

func main() {
	// Constant satuan
	const firstName string = "adid"
	const lastName = "aditya"

	// Constant Multiple
	const (
		name = "Aditya"
		age  = 24
	)

	// Constant tidak akan dikomplain oleh sistem jika tidak digunakan

	fmt.Println(name, age)
}