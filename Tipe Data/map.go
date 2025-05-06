package main

import (
	"fmt"
)

func main() {
	// bisa gini
	profile := map[string]string{
		"name": "Aditya M Zaini",
		"age": "26",
		"gender": "Male",
	}

	// bisa gini
	var shoes = map[string]string{
		"brand": "ortus",
		"color": "black",
	}

	// bisa gini
	var car = map[string]string{}
	car["brand"] = "Mercedez"
	car["color"] = "Red"

	fmt.Println(profile)
	fmt.Println(profile["name"])
	fmt.Println(shoes)
	fmt.Println(shoes["color"])
	fmt.Println(car)
	fmt.Println(car["brand"])

	// Function
	// len(map) --> Untuk mendapatkan jumlah data di map
	// map[key] --> Mengambil data di map dengan key
	// map[key] = value --> Mengubah data di map dengan key
	// make(map[TypeKey]TypeValue) --> Membuat map baru
	// delete(map, key) --> Menghapus data di map dengan key

	book := make(map[string]string)
	book["title"] = "Golang"
	book["edition"] = "1"
	book["price"] = "Rp30.000"

	fmt.Println(book)

	delete(book, "edition")

	fmt.Println(book)
}