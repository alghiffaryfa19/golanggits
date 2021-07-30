package main

import (
	"fmt"
)

func main() {

	var word string
    fmt.Print("Masukkan kata yang akan di reserve: ")
    fmt.Scanf("%s", &word)
    data := []rune(word)
    result := []rune{}
    for i := len(data) - 1; i >= 0; i-- {
        result = append(result, data[i])
    }

    fmt.Println(string(result))
}