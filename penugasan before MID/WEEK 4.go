package main

import "fmt"

func main() {
	// Masukan: harga beli, harga jual, dan jumlah saham
	var hargaBeli, hargaJual, jumlahSaham float64
	fmt.Print("Masukkan harga beli: ")
	fmt.Scan(&hargaBeli)
	fmt.Print("Masukkan harga jual: ")
	fmt.Scan(&hargaJual)
	fmt.Print("Masukkan jumlah saham: ")
	fmt.Scan(&jumlahSaham)

	// a) Menghitung total investasi awal
	totalInvestasiAwal := hargaBeli * jumlahSaham

	// b) Menghitung total penjualan
	totalPenjualan := hargaJual * jumlahSaham

	// c) Menghitung keuntungan kotor
	keuntunganKotor := totalPenjualan - totalInvestasiAwal

	// d) Menghitung biaya transaksi (0.2% dari total penjualan)
	biayaTransaksi := 0.002 * totalPenjualan

	// e) Menghitung pajak keuntungan (10% dari keuntungan kotor)
	pajakKeuntungan := 0.1 * keuntunganKotor

	// f) Menghitung keuntungan bersih
	keuntunganBersih := keuntunganKotor - biayaTransaksi - pajakKeuntungan

	// Keluaran: menampilkan keuntungan bersih
	fmt.Printf("Keuntungan Bersih: %.2f\n", keuntunganBersih)
}