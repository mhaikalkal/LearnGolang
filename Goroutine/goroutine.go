package main

import (
	"fmt"
	"runtime"
	"sort"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(3) // jumlah core hardware yang dipakai untuk eksekusi GOROUTINE

	// PROSES YANG AKAN DIJADIKAN SEBUAH GOROUTINE HARUS BERUPA FUNCTION

	// Goroutine bersifat Asynchronous
	// Yang artinya "Code akan dieksekusi tanpa menunggu eksekusi code lain selesai."
	go print(5, "Dayum")
	print(5, "Hello guys.")

	// output kedua print diatas akan random tergantung (Resource Core) / Hardware Komputer yang menentukan.
	// Bisa nabrak. kadang, for example:
	// Output: 1. Dayum 2. Dayum 1. Hello guys. 3. Dayum 2. Hello Guys. etcetc...
	// Sewaktu kita re-run ada kemungkinan ke random lagi eksekusinya

	// Scanln(v) // untuk meminta input di terminal
	// dengan memisahkan setiap parameter menggunakan spasi

	// var s1, s2, s3 string
	// fmt.Scanln(&s1, &s2, &s3)

	// user inputs: "trafalgar d law"
	// fmt.Println(s1) // trafalgar
	// fmt.Println(s2) // d
	// fmt.Println(s3) // law
}

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Printf("%v. %s\n", (i + 1), message)
	}
}

// Wait Group Sync.
func Soal_1(){
	// Soal 1
	fmt.Println("=============== Soal 1 ===============")
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	sort.Strings(phones)
	var wg sync.WaitGroup

	for i, phone := range phones{
		wg.Add(1)
		go printPhones(i+1, phone, &wg)
		time.Sleep(time.Second)
	}
	
	// menunggu func / code selesai dieksekusi. baru mulai eksekusi lagi. jadi biar ga numpuk sekaligus.
	wg.Wait() 
	fmt.Println()
}

func printPhones(row int, phone string, wg *sync.WaitGroup) {
	fmt.Printf("%v. %s\n", row, phone)
	wg.Done() // memberi tahu wg.Wait() bahwa func selesai dieksekusi.
}