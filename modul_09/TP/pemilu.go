package main

import (
	"fmt"
)

func main() {
	var umur int
	var kewarganegaraan string

	// Meminta input umur
	fmt.Print("Masukkan umur Anda: ")
	fmt.Scan(&umur)

	// Meminta input kewarganegaraan
	fmt.Print("Masukkan kewarganegaraan Anda (WNI/WNA): ")
	fmt.Scan(&kewarganegaraan)

	// Mengecek syarat untuk mengikuti pemilu
	if umur >= 17 && kewarganegaraan == "WNI" {
		fmt.Println("Anda bisa mengikuti pemilu")
	} else {
		if umur < 17 {
			fmt.Println("Anda tidak bisa mengikuti pemilu karena umur Anda belum mencapai 17 tahun.")
		}
		if kewarganegaraan != "WNI" {
			fmt.Println("Anda tidak bisa mengikuti pemilu karena Anda bukan WNI.")
		}
	}
}