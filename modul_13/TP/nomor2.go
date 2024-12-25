package main

import "fmt"

func main() {

	var angka, jumlah int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	// Perulangan untuk memeriksa faktor dari angka yang dimasukan
	for i := 1; i < angka; i++ {
		if angka%i == 0 {
			jumlah += i
		}
	}

	// Percabangan untuk memeriksa apakah benar angka tersebut merupakan angka sempurna
	if jumlah == angka {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}

}