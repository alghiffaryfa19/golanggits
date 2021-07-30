package main

import (
	"fmt"
)

func palindrome() bool {
	var word string
    fmt.Print("Masukkan kata yang akan di cek: ")
    fmt.Scanf("%s", &word)
    for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	if palindrome() == true {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan Palindrome")
	}
}