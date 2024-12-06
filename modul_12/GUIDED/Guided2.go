package main

import "fmt"

func main() {

	var angka int

	// Perulangan akan ters berjalan dan akan berhenti ketika memasukan angka lebih dari 0
	for selesai := false; !selesai; {

		fmt.Print("Masukkan angka: ")
		fmt.Scan(&angka)

		selesai = (angka > 0)

	}

	fmt.Println(angka, "adalah bilangan bulat positif")
}