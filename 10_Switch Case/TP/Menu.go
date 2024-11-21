package main

import (
	"fmt"
)

func main() {
	var menu int

	fmt.Println("1. Burger - Rp25,000", "\n2. Fried Chicken - Rp20,000", "\n3. French Fries - Rp15,000", "\n4. Soft Drink - Rp10,000", "\n5. Coffee - Rp15,000")
	fmt.Print("Masukan angka 1-5: ")
	fmt.Scan(&menu)

	switch menu {
	case 1:
		fmt.Println("Anda memilih burger - Rp25.000")
	case 2:
		fmt.Println("Anda memilih Fried Chicken - Rp20.000")
	case 3:
		fmt.Println("Anda memilih French Fries - Rp15.000")
	case 4:
		fmt.Println("Anda memilih Soft Drink - Rp10.000")
	case 5:
		fmt.Println("Anda memilih Coffe - Rp15.000")
	default:
		fmt.Println("Kode menu tidak valid")
	}
}