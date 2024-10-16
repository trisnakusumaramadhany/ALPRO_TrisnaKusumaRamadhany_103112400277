package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel untuk menyimpan jumlah baris
	var n int

	// Meminta pengguna memasukkan jumlah baris segitiga bintang
	fmt.Print("Masukkan jumlah baris segitiga bintang: ")
	fmt.Scan(&n)

	// Memeriksa apakah jumlah baris yang dimasukkan lebih dari 0
	if n > 0 {
		// Perulangan untuk mencetak segitiga bintang
		for i := 1; i <= n; i++ {
			// Perulangan untuk mencetak bintang di setiap baris
			for j := 1; j <= i; j++ {
				// Mencetak bintang tanpa pindah ke baris baru
				fmt.Print("*")
			}
			// Pindah ke baris baru setelah mencetak bintang pada satu baris
			fmt.Println()
		}
	} else {
		// Jika input tidak valid, beri pesan kesalahan
		fmt.Println("Masukkan jumlah baris yang valid (lebih dari 0)!")
	}
}