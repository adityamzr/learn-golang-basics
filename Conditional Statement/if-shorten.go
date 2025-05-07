package main

import "fmt"

func main() {
	name := "salsabila"

	if length := len(name); length > 5 {
		fmt.Println("Nama sudah bagus")
	}else{
		fmt.Println("Nama terlalu pendek")
	}
}