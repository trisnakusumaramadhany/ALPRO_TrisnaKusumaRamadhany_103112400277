package main

import (
	"fmt"
	"strings"
)

func main() {
	var huruf string

	// Meminta input dari pengguna
	fmt.Print("Masukkan sebuah huruf: ")
	fmt.Scan(&huruf)

	// Mengubah huruf menjadi kapital untuk memudahkan pengecekan
	huruf = strings.ToUpper(huruf)

	// Mengecek apakah inputan adalah huruf vokal
	if len(huruf) == 1 && strings.ContainsAny(huruf, "AEIOU") {
		fmt.Println("Huruf Vokal")
	} else if len(huruf) == 1 && strings.ContainsAny(huruf, "BCDFGHJKLMNPQRSTVWXYZ") {
		fmt.Println("Huruf Konsonan")
	} else {
		fmt.Println("Input tidak valid. Harap masukkan satu huruf.")
	}
}