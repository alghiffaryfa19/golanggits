package main

import (
	"fmt"
	"strings"
)

func main() {
	var email string
	fmt.Print("Masukkan email: ")
    fmt.Scanf("%s", &email)
	if strings.Contains(email, "@") {
		components := strings.Split(email, "@")
		domainname := components[1]
		if strings.Contains(domainname, ".") {
			user := components[0]
			if len(user) > 20 {
				fmt.Printf("Email tidak valid, karena username lebih dari 20 karakter")
			} else {
				item := strings.SplitAfterN(domainname, ".", 2)
				domain := item[1]
				if domain == "co.id" || domain == "id" {
					fmt.Printf("Email valid, memenuhi semua syarat")
				} else {
					fmt.Printf("Email tidak valid, karena domain tidak co.id atau id")
				}
			}
		} else {
			fmt.Printf("Email tidak valid, karena tidak ada . setelah @")
		}
	} else {
		fmt.Printf("Email tidak valid, karena tidak ada @")
	}
}