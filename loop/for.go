package main

import "fmt"

func main()  {
	fmt.Println("Mulai loop...");
	// Cara 1
	// counter := 1

	// for counter <= 5 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	Cara 2
	for counter := 1; counter <= 5; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	fmt.Println("Selesai");
}