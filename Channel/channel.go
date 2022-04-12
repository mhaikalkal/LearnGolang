package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

/*
	Channel digunakan untuk menghubungkan goroutine satu dengan goroutine lain.
	Mekanisme yang dilakukan adalah serah-terima data lewat channel tersebut.
	Dalam komunikasinya, sebuah channel difungsikan sebagai pengirim di sebuah goroutine, dan juga sebagai penerima di goroutine lainnya.
	Pengiriman dan penerimaan data pada channel bersifat blocking atau synchronous.
*/

func main() {
	// channelAngka()
	// channelHello()
	// channelGoIIFE()
	// channelSlice()

	// Close
	// channelClose()
	// Soal_2()

	// channelBuffer()

	// select
	// channelSelect()
	// Soal_3()

	// channelWaitGrup()

	// channelForRange()

	channelTimeOut()
	
}

func channelAngka(){
	// declare channel: var varName = make(chan typeData)
	// untuk mengisi channel, harus menggunakan arrow <-
	var angka = make(chan interface{})

	var printNumber = func(num chan interface{}) {
		num <- 10 // assign var angka/(num), menggunakan channel.
	}

	go printNumber(angka)

	// menerima nilai dari channel
	nilai := <- angka
	fmt.Println("var Nilai menerima angka dari channel var Angka : ", nilai)
}


func channelHello(){
	// say Hello
	var msg = make(chan string)

	go sayHelloTo("Haikal", &msg)
	go sayHelloTo("Peter", &msg)

	var message1 = <- msg
	fmt.Println(message1) // nge print Haikal

	var message2 = <- msg
	fmt.Println(message2) // nge print Peter
}

func sayHelloTo(people string, msg *chan string) {
	var data = fmt.Sprintf("Hello, %s", people)
	*msg <- data
}

// Eksekusi Goroutine pada IIFE
func channelGoIIFE(){
	var messages = make(chan string)

	go func(who string) {
		var data = fmt.Sprintf("Dayum, %s. You looking fine asf.", who)
		messages <- data
	}("Haikal")

	var mess = <-messages
	fmt.Println(mess)
}

func channelSlice() {
	var numbChannel = make(chan int)
	var sliceNumb = []int{69, 77, 21, 11}

	for _, eachNumber := range sliceNumb{
		go func(angka int) { // pake iife func
			numbChannel <- eachNumber // dibaca: setiap number kita jadikan sebuah channel numbChannel
		}(eachNumber) // untuk mengisi (angka int) iife & anonym parameter func
	}

	// Print
	for i := 0; i < len(sliceNumb); i++{
		fmt.Println(<-numbChannel)
		time.Sleep(time.Second)
	}
}


// Close channel
func channelClose() {
	ch := make(chan int)
  
	go cetak(ch)
	
	for {
		data, ok := <-ch // dari fungsi close(ch)
		if ok == false {
			break
		}

		fmt.Printf("Data di terima %v\n", data)
	}
}

func cetak(ch chan<- int) { // ch merupakan chan yg menerima int. ch chan<- int
	for index := 0; index < 10; index++ {
	  ch <- index // index diassing ke ch. berarti 0,1,2,...,9
	}
	close(ch) // Selesai.
	// Jadi kalau memasukkan value ke channel. memakai Looping
	// harus di close
}

// Close Channel Soal 2
func Soal_2(){
	fmt.Println("=============== Soal 2 ===============")
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}
	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	fmt.Println("List Movies :")
	for value := range moviesChannel {
		fmt.Println(value)
	}

	fmt.Println()
}

func getMovies(movCH chan string, movies...string) {
	for i, mov := range movies{
		movCH <- fmt.Sprintf("%v. %s", i+1, mov)
	}

	// ketika kita memasukkan data memakai loop ke sebuah channel. 
	// kita harus memberi tahu bahwa 
	close(movCH)
}


// Buffer, antrian
func channelBuffer(){
	// namaChannel := make(chan TypeChannel, capacityChannel)
	runtime.GOMAXPROCS(2)
	ch := make(chan int, 3)

	ch <- 6
	ch <- 7
	ch <- 5
	// ch <- 69

	// Jika lebih dari capacityChannel. dia akan deadlock
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch)

	fmt.Println("========================================")
	alert := make(chan int, 3)

	go func() {
		for {
			i := <- alert
			fmt.Println("Menerima data:", i)
		}
	}() // self run

	for i := 1; i <= 5; i++{
		fmt.Println("Send Data", i)
		alert <- i // masukin i ke channel alert
		// time.Sleep(time.Second)
	}

}

// Penggunaan Select
func getAverage(numbers []int, ch chan float64) {
    var sum = 0
    for _, e := range numbers {
        sum += e
    }
    ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
    var max = numbers[0]
    for _, e := range numbers {
        if max < e {
            max = e
        }
    }
    ch <- max
}

func channelSelect(){
	runtime.GOMAXPROCS(2)

    var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
    fmt.Println("numbers :", numbers)

    var ch1 = make(chan float64)
    go getAverage(numbers, ch1)

    var ch2 = make(chan int)
    go getMax(numbers, ch2)

    for i := 0; i < 2; i++ {
        select {
        case avg := <-ch1:
            fmt.Printf("Avg \t: %.2f \n", avg)
        case max := <-ch2:
            fmt.Printf("Max \t: %d \n", max)
        }
    }
}

// Select Soal 3
func Soal_4(){
	fmt.Println("=============== Soal 4 ===============")
	var bal = balok{
		persegiPanjang: persegiPanjang{panjang: 10, lebar: 5},
		tinggi: 20,
	}

	var luas = make(chan interface{})
	var keliling = make(chan interface{})
	var volume = make(chan interface{})

	go luasPersegi(bal, &luas)
	go kelilingPersegi(bal, &keliling)
	go volumeBalok(bal, &volume)

	for i := 1; i <= 3; i++{
		select{
		case luasPersegi := <- luas:
			fmt.Println("Luas Persegi dengan panjang", bal.panjang, "dan lebar", bal.lebar, "adalah :", luasPersegi)
		case kelilingPersegi := <- keliling:
			fmt.Println("Keliling Persegi dengan panjang", bal.panjang, "dan lebar", bal.lebar, "adalah :", kelilingPersegi)
		case volumeBalok := <- volume:
			fmt.Println("Volume Balok dengan panjang", bal.panjang, "lebar", bal.lebar, "dan tinggi", bal.tinggi, "adalah :", volumeBalok)
		}
	}
}

type persegiPanjang struct{
	panjang int
	lebar int
}

type balok struct{
	persegiPanjang
	tinggi int
}

func luasPersegi(b balok, luas *chan interface{}){
	*luas <- b.lebar * b.panjang
}

func kelilingPersegi(b balok, keliling *chan interface{}){
	*keliling <- 2 * (b.lebar + b.panjang)
}

func volumeBalok(b balok, volume *chan interface{}){
	*volume <- b.lebar * b.panjang * b.tinggi
}


// channel wait grup
func printText(text string, wg *sync.WaitGroup){
	for i:=0; i<5; i++{
		fmt.Println(text)
	}
	wg.Done()
}

func channelWaitGrup(){
	var wg sync.WaitGroup

	wg.Add(1) // menambahkan proses goRoutine yang berjalan.
	go printText("Halo", &wg)

	wg.Add(1)
	go printText("Dunia", &wg)

	wg.Wait()
}

// Channel DI FOR - RANGE
func channelForRange(){
	var messages = make(chan string)
	go sendMessage(messages)
	printMessage(messages)
}

func sendMessage(ch chan<- string) {
    for i := 0; i < 20; i++ {
        ch <- fmt.Sprintf("data %d", i)
    }
    close(ch)
}

func printMessage(ch <-chan string) {
    for message := range ch {
        fmt.Println(message)
    }
}

/*
	||		Syntax		||							USE FOR								||
	|| ch chan string	|| Parameter ch bisa digunakan untuk mengirim dan menerima data	||
	|| ch chan<- string	|| Parameter ch hanya bisa digunakan untuk mengirim data		||
	|| ch <-chan string || Parameter ch hanya bisa digunakan untuk menerima data		||
*/

// Channel Timeout
func channelTimeOut(){
	rand.Seed(time.Now().Unix())
    runtime.GOMAXPROCS(2)

	var msg = make(chan interface{})

	go sendData(msg)
	retrieveData(msg)	
}

func sendData(ch chan<- interface{}){
	for i := 1; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int() % 10 + 1) * time.Second)
	}
}

func retrieveData(msg chan interface{}){
	loop:
	for {
		select{
		case data := <- msg:
			fmt.Print(`receive data "`, data, `"`, "\n")
		case <-time.After(time.Second * 5):
			fmt.Println("Conection Timeout. No Activities under 5 seconds")
			break loop // ulang lagi ke atas
		}

	}

}