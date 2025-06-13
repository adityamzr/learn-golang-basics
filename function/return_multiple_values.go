package main

import "fmt"

func getName() (string, string) {
	return "aditya", "Mzr"
}

func main() {
	// firstName, lastName := getName()
	firstName, _ := getName()
	fmt.Println(firstName)
}