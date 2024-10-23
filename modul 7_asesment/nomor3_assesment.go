package main

import (
	"fmt"
)

func main() {
	var qirat int
	fmt.Print("Masukkan jumlah uang dalam satuan qirat: ")
	fmt.Scan(&qirat)

	// Konversi dari qirat ke fals
	fals := qirat / 6
	// Sisa qirat setelah konversi ke fals
	sisaQirat := qirat % 6

	// Konversi dari fals ke dirham
	dirham := fals / 10
	// Sisa fals setelah konversi ke dirham
	sisaFals := fals % 10

	// Konversi dari dirham ke dinar
	dinar := dirham / 10
	// Sisa dirham setelah konversi ke dinar
	sisaDirham := dirham % 10

	// Menampilkan hasil konversi
	fmt.Printf("Hasil konversi:\n")
	fmt.Printf("Dinar: %d\n", dinar)
	fmt.Printf("Dirham: %d\n", sisaDirham)
	fmt.Printf("Fals: %d\n", sisaFals)
	fmt.Printf("Qirat: %d\n", sisaQirat)
}