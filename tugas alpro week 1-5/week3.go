/*
Soal nomer 4 terurut 2
Sebuah program digunakan untuk menentukan suatu bilangan memiliki digit terurut atau tidak Masukan terdiri dari tiga digit bilangan bulat.
Keluaran berupa boolean yang menyatakan setiap digit pada bilangan bulat yang diberikan terurut atau tidak

#pseucode

program bilangan
kamus
    bilangan: integer
    status: boolean

algoritma
    input(bilangan)
    status bilangan akan di cek dengan cara di bagi 10 dan di cari modulus dari 10 dan akan di cek dengan kondisi < atau >
    output (status)

endprogram
*/

package main

import "fmt"

func main() {
	var bilangan int
	var status bool
	fmt.Scan(&bilangan)
	status = (bilangan/100 < (bilangan/10)%10) && ((bilangan/10)%10 < bilangan%10) || (bilangan/100 > (bilangan/10)%10 && (bilangan/10)%10 > bilangan%10)
	fmt.Println(status)
}
