package main

import "fmt"

func main() {

	var angka, jumlah int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	// Perulangan yang digunakan untuk menghitung jumlah digit angka
	for selesai := false; !selesai; {

		angka = angka / 10
		jumlah++

		selesai = (angka == 0)
	}

	fmt.Print(jumlah)

}