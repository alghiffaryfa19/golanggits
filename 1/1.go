package main

import (
	"fmt"
)

func main() {

	var number int

	fmt.Println("Masukkan Angka : ")
	fmt.Scan(&number)

    if number%5 == 0 && number%3 == 0 {
        fmt.Println("Hello World")
    } else if number%3 == 0 {
        fmt.Println("Hello")
    } else if number%5 == 0  {
		fmt.Println("World")
	} else {
		fmt.Println("Bukan angka yang habis dibagi 3 dan/atau 5")
	}
}