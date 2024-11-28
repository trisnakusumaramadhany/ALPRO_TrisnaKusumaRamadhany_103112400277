package main

import (
	"fmt"
)

func main() {
	var total float64
	for { // While-loop simulasi
		var item string
		fmt.Print("Masukkan nama barang (atau ketik 'selesai' untuk keluar): ")
		fmt.Scanln(&item)

		if item == "selesai" {
			break
		}

		var price float64
		fmt.Print("Masukkan harga barang: ")
		fmt.Scanln(&price)

		total += price
	}

	fmt.Printf("Total belanja: %.2f\n", total)
	fmt.Println("Transaksi selesai. Terima kasih!")
}