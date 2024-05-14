package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Nava"
	names[1] = "Dwi"
	names[2] = "Indriyanti"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var val = [3]int{
		23, 24, 43,
	}

	fmt.Println(val)

	// len itu bukan jumlah dari data array melainkan panjang dari si array nya
	fmt.Println(len(names))
}