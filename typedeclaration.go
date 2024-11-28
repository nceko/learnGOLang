package main

import "fmt"

func main() {
	type (
		NIB       int32
		namaOrang string
		status    bool
	)

	var (
		pegawai       NIB       = 101359
		namaPegawai   namaOrang = "Eko Hidayat"
		statusPegawai status    = true
	)

	fmt.Println(pegawai)
	fmt.Println(namaPegawai)
	fmt.Println(statusPegawai)
}
