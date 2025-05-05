package main

import "fmt"

func main() {
	var (
		nilai32 int32 = 3211
		nilai64 int64 = int64(nilai32)
		nilai8  int8  = int8(nilai32)
	)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8) /* Karna melebihi kapasitas tipe data maka hasilnya akan berbeda */

	var (
		name = "did"
		i = name[0]
		iString = string(i)
	)

	fmt.Println(iString)

}