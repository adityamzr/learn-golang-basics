package main

import "fmt"

func main() {
	name := "adid"

	// IF STATEMENT
	if name == "adid" {
		fmt.Println(name)
	}

	// IF ELSE STATEMENT
	if name == "adid" {
		fmt.Println("Nama kamu adid")
	}else {
		fmt.Println("Siapa nama kamu?")
	}

	// ELSE IF STATEMENT
	if name == "adid" {
		fmt.Println("Hi adid!")
	}else if name == "zain" {
		fmt.Println("Hi zain!")
	}else {
		fmt.Println("Nama kamu siapa?")
	}

	// SHORTEN IF STATEMENT
	if length := len(name); length > 5 {
		fmt.Println("Nama kamu panjang ya")
	}else {
		fmt.Println("Nama kamu pendek tapi bagus")
	}
}