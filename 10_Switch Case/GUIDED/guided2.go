package main

import "fmt"

func main() {

	// Membuat variable tanaman dan inputan untuk memasukan nama tanaman
	var nama_tanaman string
	fmt.Scan(&nama_tanaman)

	// Switch case untuk menentukan tanaman karnivora
	switch nama_tanaman {
	case "nepenthes", "drosera":
		fmt.Println("Termasuk Tanaman Karnivora.")
		fmt.Println("Asli Indonesia.")
	case "venus", "sarracenia":
		fmt.Println("Termasuk Tanaman Karnivora.")
		fmt.Println("Tidak Asli Indonesia.")
	// default digunakan untuk menampilkan teks jika tanaman tidak termasuk karnivora
	default:
		fmt.Println("Tidak termasuk Tanaman Karnivora.")
	}
}