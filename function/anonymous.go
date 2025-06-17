package main

import (
	"fmt"
	"strings"
)

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(strings.ToLower(name)) {
		fmt.Println(name, "is not allowed!")
	}else {
		fmt.Println(name, "is allowed!")
	}
}

func main()  {
	blacklist := func(name string) bool {
		return name == "bodoh"
	}
	registerUser("Adid", blacklist)
	// atau
	registerUser("Kumar", func(name string) bool {
		return name == "bodoh"	
	})
}