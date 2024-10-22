package main

import "fmt"

func main() {
	var jumlahBarang int
	fmt.Print("Masukkan jumlah barang yang dibeli: ")
	fmt.Scanln(&jumlahBarang)

	var totalPoin int = 0

	// Menghitung poin dasar
	totalPoin = jumlahBarang * 10

	// Menghitung poin tambahan
	if jumlahBarang > 5 {
		totalPoin += (jumlahBarang - 5) * 5
	}

	fmt.Printf("Total poin yang didapatkan: %d poin\n", totalPoin)
}