package main

import (
	"fmt"
)

func main() {
	var umur int

	fmt.Print("Masukan umur 1-100: ")
	fmt.Scan(&umur)

	switch {
	case umur >= 0 && umur <= 12:
		fmt.Println("Usia Anda: ", umur)
		fmt.Println("Kategori: Anak-Anak")
	case umur >= 13 && umur <= 17:
		fmt.Println("Usia Anda: ", umur)
		fmt.Println("Kategori: Remaja")
	case umur >= 18 && umur <= 64:
		fmt.Println("Usia Anda: ", umur)
		fmt.Println("Kategori: Dewasa")
	case umur > 64:
		fmt.Println("Usia Anda: ", umur)
		fmt.Println("Kategori: Lansia")

	}
}