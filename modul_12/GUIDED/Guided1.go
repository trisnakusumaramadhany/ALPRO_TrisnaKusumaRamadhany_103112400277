package main

import "fmt"

func main() {

	var kata string
	var jumlah int

	fmt.Print("Masukkan kata: ")
	fmt.Scan(&kata)

	fmt.Print("Masukkan jumlah: ")
	fmt.Scan(&jumlah)

	trigger := 0

	// Perulangan akan mengulang kata yang dimasukan sesuai dengan jumlah yang dimasukan juga.
	for selesai := false; !selesai; {
		fmt.Println(kata)
		trigger++

		selesai = (trigger >= jumlah)
	}
}