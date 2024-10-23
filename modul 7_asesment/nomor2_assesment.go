package main

import (
	"fmt"
)

func main() {
	var N int

	// Meminta pengguna untuk memasukkan bilangan bulat positif N
	fmt.Print("Masukkan bilangan bulat positif N: ")
	_, err := fmt.Scan(&N)

	// Memeriksa apakah input valid
	if err != nil || N <= 0 {
		fmt.Println("Input tidak valid. Harap masukkan bilangan bulat positif.")
		return
	}

	// Mencetak kuadrat dari bilangan 1 sampai N
	fmt.Printf("Kuadrat dari bilangan 1 sampai %d:\n", N)
	for i := 1; i <= N; i++ {
		fmt.Printf("%d kuadrat = %d\n", i, i*i)
	}
}