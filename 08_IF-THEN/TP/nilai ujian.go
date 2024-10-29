package main

import (
	"fmt"
)

func main() {

	// Membuat variabel
	var nilai int

	// Inputan untuk memasukan nilai ujian
	fmt.Print("Masukan nilai ujian anda: ")
	fmt.Scan(&nilai)

	// If else yang dimana digunakan untuk melakukan operasi, yaitu ketika nilai yang dimasukan lebih dari sama dengan 70, maka akan ditampilkan Lulus, jika kurang dari sama dengan 70, maka akan muncul tidak lulus.
	if nilai >= 70 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Tidak lulus")
	}
}