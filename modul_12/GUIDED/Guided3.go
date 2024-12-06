package main

import "fmt"

func main() {

	var angka, kelipatan int
	var selesai bool

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	fmt.Print("Masukkan kelipatan: ")
	fmt.Scan(&kelipatan)

	for selesai = false; !selesai; {

		angka = angka - kelipatan

		fmt.Println(angka)
		selesai = (angka <= 0)

	}

	fmt.Println(angka == 0)
}