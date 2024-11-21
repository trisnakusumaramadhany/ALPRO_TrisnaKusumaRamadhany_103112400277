package main

import "fmt"

func main() {

	// Membuat variable jam24 untuk menyimpan masukan, jam12 untuk menyimpan hasil, label untuk menyimpan label
	var jam24, jam12 int
	var label string

	// Meminta untuk memasukan jam
	fmt.Print("Masukkan jam (0-23): ")
	fmt.Scan(&jam24)

	// Switch case untuk menentukan jam yang dimasukan apakah termasuk AM atau PM
	switch {
	case jam24 >= 0 && jam24 <= 11:
		label = "AM"
		jam12 = jam24
	case jam24 >= 12 && jam24 <= 23:
		label = "PM"
		jam12 = jam24 - 12
	default:
		jam12 = jam12
		label = " Bukan termasuk jam"
	}

	// Menampilkan hasil switch case sesuai dengan kondisi yang terpenuhii.
	fmt.Print(jam12, label)

}