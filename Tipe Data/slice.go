package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = months[10:]
	// Ketika append tidak merubah array sebelumnya, karena membuat array baru dan menambah data
	var slice3 = append(slice2, "Ramadhan")
	
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)
}