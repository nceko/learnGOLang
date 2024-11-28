package main

import "fmt"

func main() {
	var name = "Eko Hidayat"
	fmt.Println(name)

	name = "Hidayat"
	fmt.Println(name)

	// PENGGUNAAN TANPA VAR
	test := "nama saya eko hidayat"
	fmt.Println(test)

	// PENGGUNAAN MULTIPLE VARIABLE
	// var firstName, lastName = "Eko", "Hidayat"

	// fmt.Println(firstName)
	// fmt.Println(lastName)

	var (
		firstName = "Eko"
		lastName  = "Hidayat"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
