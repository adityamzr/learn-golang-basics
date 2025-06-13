package main

import "fmt"

func main()  {
	// Mengiterasi data
	names := []string{
		"Aditya",
		"Muhammad",
		"Zaini",
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// Menggunakan for range
	for index, name := range names {
		fmt.Println("Index ke", index, "=", name)
	}

	// Menggunakan for range tanpa index
	for _, name := range names {
		fmt.Println(name)
	}
}