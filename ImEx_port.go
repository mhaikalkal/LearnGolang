package main

import "fmt"

func SayHello(name string) { // Kalau nama func awalnya pake Huruf besar maka scope nya Global. 
	fmt.Printf("Hello")
	introduce(name) // 
}

// kalau huruf awal func-nya huruf kecil maka scope nya private / local.
// jadi kalau misalkan file ini beda package dan pengen dipanggil di file main.go pasti error.
// jadi biar bisa masuk ke file main.go, kita masukin ke func yang scope nya global.
func introduce(name string){  
	fmt.Printf(". Nama saya adalah %s.\n", name)
}

func calculate(nums...int) int {
	var sum int
	for _, num := range nums{
		sum += num
	}

	return sum
}