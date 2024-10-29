package main

import (
	"fmt"
)

func main() {

	// Variabel untuk menyimpan nilai dari angka yang dimasukan
	var angka int

	// Inputan untuk memasukan angka
	fmt.Print("Masukan angka: ")
	fmt.Scan(&angka)

	// Percabangan untuk menentukan apakah angka tersebut genap atau ganjil (Ketika angka habis dibagi 2, angka tersebut adalah angka genap, sebaliknya jika sisa maka angka tersebut merupakan ganjil)
	if angka%2 == 0 {
		fmt.Println("Angka adalah Genap")
	} else {
		fmt.Println("Angka adalah Ganjil")
	}
}