package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Aditya"
	middleName = "M"
	lastName = "Zaini"

	return firstName, middleName, lastName
}

func main()  {
	x, y, z := getFullName()
	fmt.Println(x, y, z)
}