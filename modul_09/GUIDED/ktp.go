package main 

import "fmt"

func main (){
	var usia int
	var KK bool

	fmt.Print("masukan Usia dan KK anda")
	fmt.Scan(&usia,&KK)

	if usia >= 17 && KK == true {
		fmt.Print("bisa membuat KTP")
	
	} else {
		fmt.Print("belum bisa membuat ktp")
	}



}