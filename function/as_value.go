package main

import "fmt"

func wasap(name string) string  {
	return "Wazaaapp mas " + name
}

func main()  {
	tes := wasap
	tos := wasap("jon")
	
	fmt.Println(tes("adid"))
	fmt.Println(tos)
}