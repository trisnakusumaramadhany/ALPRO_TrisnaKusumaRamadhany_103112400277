package main

import (
	"fmt"
)

func main() {
	var number int

	// Meminta pengguna memasukkan bilangan 4 digit
	fmt.Print("Masukkan bilangan 4 digit (1000 - 9999): ")
	fmt.Scan(&number)

	// Memeriksa apakah bilangan memenuhi syarat 4 digit
	if number < 1000 || number > 9999 {
		fmt.Println("Bilangan harus 4 digit dan berada antara 1000 hingga 9999.")
		return
	}

	// Mengambil setiap digit dari bilangan
	digit1 := number / 1000            // ribuan
	digit2 := (number / 100) % 10      // ratusan
	digit3 := (number / 10) % 10       // puluhan
	digit4 := number % 10              // satuan

	// Mengecek pola urutan digit
	if digit1 < digit2 && digit2 < digit3 && digit3 < digit4 {
		fmt.Println("Digit terurut membesar")
	} else if digit1 > digit2 && digit2 > digit3 && digit3 > digit4 {
		fmt.Println("Digit terurut mengecil")
	} else {
		fmt.Println("Digit tidak terurut")
	}
}