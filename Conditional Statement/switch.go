package main

import "fmt"

func main() {
	name := "salsa"

	// NORMAL SWITCH
	switch name {
	case "salsabila":
		fmt.Println("Hi salsa!")
	case "syah":
		fmt.Println("Hi syah!")
	default:
		fmt.Println("Nama kamu siapa?")
	}

	// SHORTEN SWITCH
	switch length := len(name); length > 5 {
	case true: fmt.Println("Nama lebih dari 5 huruf")
	case false: fmt.Println("Nama kurang dari 5 huruf")
	}

	// SWITCH WITHOUT CONDITION - SIMILAR TO ELSE IF
	length:= len(name)
	switch {
	case length > 5: fmt.Println("Nama lebih dari 5 huruf")
	case length < 5: fmt.Println("Nama kurang dari 5 huruf")
	default: fmt.Println("Nama 5 huruf")
	}
}