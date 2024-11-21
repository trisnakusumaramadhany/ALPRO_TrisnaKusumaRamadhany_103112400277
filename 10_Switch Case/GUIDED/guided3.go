package main

import "fmt"

func main() {

	// Membuat variabel untuk menyimpan jenis kendaraan, durasi dan tarif
	var kendaraan string
	var durasi int
	var tarif int

	// Inputan untuk memasukan jenis kendaraan dan durasi
	fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	// Switch case untuk menentukan tarif sesuai dengan jenis kendaraan dan durasi
	switch {
	case kendaraan == "Motor" && durasi >= 1 && durasi <= 2:
		tarif = 7000
	case kendaraan == "Motor" && durasi > 2:
		tarif = 9000
	case kendaraan == "Mobil" && durasi >= 1 && durasi <= 2:
		tarif = 15000
	case kendaraan == "Mobil" && durasi > 2:
		tarif = 20000
	case kendaraan == "Truk" && durasi >= 1 && durasi <= 2:
		tarif = 25000
	case kendaraan == "Truk" && durasi > 2:
		tarif = 35000
	default:
		fmt.Println("Jenis kendaraan atau durasi parkir tidak valid")
	}

	// Menampilkan hasil
	fmt.Printf("Tarif Parkir: Rp %d\n", tarif)
}