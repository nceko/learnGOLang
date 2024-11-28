package main

import "fmt"

func main() {
	var (
		nilai32 int32 = 25468
		nilai64 int64 = int64(nilai32)
		nilai8  int8  = int8(nilai32)

		name  = "Eko Hidayat"
		char  = name[8]
		huruf = string(char)
	)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	fmt.Println(name)
	fmt.Println(char)
	fmt.Println(huruf)

}
