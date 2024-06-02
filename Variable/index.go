package main

import "fmt"

func main()  {
	//Contoh variable yang didefinisikan tipe datanya
	var name string

	name = "karina"
	fmt.Println(name)
	
	name = "karin"
	fmt.Println(name)

	//Contoh Variable yang tidak didefinisikan tipe datanya
	var myName = "Aditya"
	fmt.Println(myName)

	// Secara default variable int akan bertipe data int saja
	var age = 24
	fmt.Println(age)

	// Penggunaan var tidak wajib asalkan saat mendeklarasikan variable sudah dengan datanya
	country := "Indonesia"
	fmt.Println(country)

	// variable bisa dibuat multiple
	var (
		firstName = "Aditya"
		lastName = "Ramdani"
	) 

	fmt.Println(firstName, lastName)
} 