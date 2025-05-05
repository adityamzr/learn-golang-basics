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

	var slice1 = months[7:12]
	// fmt.Println(len(months))
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice1[0] = "Agus gustino"
	var slice2 = months[10:]
	// Ketika append tidak merubah array sebelumnya, karena membuat array baru dan menambah data
	var slice3 = append(slice2, "Ramadhan")
	
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "adid"
	newSlice[1] = "zain"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	var newSlice2 = append(newSlice, "ramdani")

	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[1] = "muhammad"

	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// COPY SLICE
	fromSlice := months[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// Perbedaan array dan slice
	iniArray := [...]int{1,2,3}
	iniSlice := []int{1,2,3}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	// Slice lebih sering digunakan daripada array di golang
}