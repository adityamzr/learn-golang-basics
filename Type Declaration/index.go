package main

import "fmt"

func main() {
	type noKtp string

	var nik noKtp = "320412414214941"

	var contoh = "19237172631"
	var gantiContoh = noKtp(contoh)

	fmt.Println(nik)
	fmt.Println(contoh)
	fmt.Println(gantiContoh)
}