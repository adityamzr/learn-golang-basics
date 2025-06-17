package main

import (
	"fmt"
	"strings"
)

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello ", filteredName)
}

func spamFilter(name string) string {
	spam := strings.ToLower(name)
	if(spam == "tolol"){
		return "*cencored*"
	}else {
		return name
	}
}

func main()  {
	sayHelloWithFilter("ADID", spamFilter)
	// atau
	filtered := spamFilter
	sayHelloWithFilter("tolol", filtered)
}