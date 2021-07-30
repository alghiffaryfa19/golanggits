package main

import (
	"fmt"
	"strconv"
)

func waktu() string {
	var time string
	fmt.Print("Masukkan waktu 12h (ex 12:00:00AM) : ")
    fmt.Scanf("%s", &time)
	if string(time[len(time)-2:]) == "AM" && string(time[:2]) == "12" {
		return "00" + string(time[2:len(time)-2])
	} else if string(time[len(time)-2:]) == "AM" {
		return string(time[0:len(time)-2])
	} else if string(time[len(time)-2:]) == "PM" && string(time[:2]) == "12" {
		timer := string(time[:2])
		intVar, err := strconv.Atoi(timer)
		if err != nil {
			fmt.Printf("Eror")
		}

		return  strconv.Itoa(intVar+12) + time[2:len(time)-2]
	} else {
		timer := string(time[:2])
		intVar, err := strconv.Atoi(timer)
		if err != nil {
			fmt.Printf("Eror")
		}

		return  strconv.Itoa(intVar+12) + time[2:len(time)-2]
	}
}

func main() {
	fmt.Printf(waktu())
}