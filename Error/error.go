package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	/*
		Di Go, banyak sekali fungsi yang mengembalikan nilai balik lebih dari satu.
		Biasanya, salah satu kembalian adalah bertipe error.
		Contohnya seperti pada fungsi strconv.Atoi().
		Fungsi tersebut digunakan untuk konversi data string menjadi numerik. Fungsi ini mengembalikan 2 nilai balik.
		Nilai balik pertama adalah hasil konversi, dan nilai balik kedua adalah error.
		Ketika konversi berjalan mulus, nilai balik kedua akan bernilai nil.
		Sedangkan ketika konversi gagal, penyebabnya bisa langsung diketahui dari error yang dikembalikan.
	*/

	var input string
	fmt.Print("Type some number: ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}

	// String validate error
	var name string
    fmt.Print("Type your name: ")
    fmt.Scanln(&name)

    if valid, err := validate(name); valid {
        fmt.Println("halo", name)
    } else {
        fmt.Println(err.Error())
    }
}

func validate(input string) (bool, error) {
    if strings.TrimSpace(input) == "" {
        return false, errors.New("cannot be empty")
    }
    return true, nil
}