package main

import (
    "fmt"
)

func main() {
    var x, y int

    // Meminta pengguna untuk memasukkan nilai x dan y
    fmt.Print("Masukkan nilai x (bilangan bulat positif): ")
    fmt.Scan(&x)
    fmt.Print("Masukkan nilai y (bilangan bulat positif, x <= y): ")
    fmt.Scan(&y)

    // Memeriksa apakah x lebih kecil atau sama dengan y
    if x > y {
        fmt.Println("Nilai x harus lebih kecil atau sama dengan y.")
        return
    }

    // Menghitung jumlah dari x sampai y
    sum := 0
    for i := x; i <= y; i++ {
        sum += i
    }

    // Menampilkan hasil
    fmt.Printf("Jumlah bilangan dari %d sampai %d adalah: %d\n", x, y, sum)
}