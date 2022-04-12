package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// Normal Interface
	interfaceFunc()

	// Embedded Interface
	embedInterfaceFunc()

	// Interface Kosong. interface{}
	emptyInterface()

	// Soal Tugas ke-8 Golang CodingCamp
	soal_1()
	soal_2()
	soal_3()
	soal_4()

}

// Interface type calculate
// Interface merupakan kumpulan definisi method tanpa isi. default value nya null/nil.
// jika ingin menggunakan sebuah interface, penamaan method sebuah struct harus sama.
type calculate interface {
	luas() float64 // luas
	keliling() float64
}

// Struct Lingkaran
type lingkaran struct {
	diameter float64
}

// method Lingkaran
func (l lingkaran) jari() float64 {
	return l.diameter / 2
}

// jika ingin menggunakan interface calculate, kalau kita menamai method ini luasLingkaran akan error
func (l lingkaran) luas() float64 { 
	return math.Pi * math.Pow(l.diameter/2, 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

// Struct Persegi
type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func interfaceFunc() {
	var bangunDatar calculate

	bangunDatar = lingkaran{7.0}
	// akses interface luasnya. karena memiliki kesamaan nama method, dia otomatis mereference method lingkarannya
	// Jika ingin mengakses property lingkaran via setelah menggunakan interface. Kita harus casting. Caranya : 
	// bangunDatar.(lingkaran).diameter. Kita panggil nama struct menggunakan bracker dan seperti memanggil property biasa.
	fmt.Printf("Luas Lingkaran yang berdiameter %v adalah : %.2f\n", bangunDatar.(lingkaran).diameter, bangunDatar.luas())
	fmt.Printf("Keliling Lingkaran yang berdiameter %v adalah : %.2f\n", bangunDatar.(lingkaran).diameter, bangunDatar.keliling())
	// karena interface bangunDatar tidak memiliki method jari. Maka kita harus akses structnya dan exec method-nya seperti biasa.
	fmt.Printf("Jari-jari dari Lingkaran yang berdiameter %v adalah : %.2f\n", bangunDatar.(lingkaran).diameter, bangunDatar.(lingkaran).jari())
		
	fmt.Println("===============================================================")
	
	bangunDatar = persegi{10}
	fmt.Printf("Luas Persegi yang memiliki sisi %v adalah : %.2f\n", bangunDatar.(persegi).sisi, bangunDatar.luas())
	fmt.Printf("Keliling Persegi yang memiliki sisi %v adalah : %.2f\n", bangunDatar.(persegi).sisi, bangunDatar.keliling())

	fmt.Println("===============================================================")
}

// Embedded Interface
type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung interface {
	hitung2d
	hitung3d
}

// struct kubus
type kubus struct {
	sisi float64
}

// method kubus, Coba pake pointer
func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

func embedInterfaceFunc() {
	// deklarasi nya gini juga bisa.
	var calculate hitung = &kubus{9}

	// karena kita reference &kubus. maka panggil valuenya pake *kubus
	fmt.Printf("Volume kubus dengan sisi %v adalah : %v\n", calculate.(*kubus).sisi, calculate.volume())
	fmt.Printf("Luas kubus dengan sisi %v adalah : %v\n", calculate.(*kubus).sisi, calculate.luas())
	fmt.Printf("Keliling kubus dengan sisi %v adalah : %v\n", calculate.(*kubus).sisi, calculate.keliling())

	fmt.Println("===============================================================")
}

// interface kosong
func emptyInterface() {
	// interface dan interface{} sangatlah berbeda.
	// interface merupakan kumpulan method.
	// sedangkan interface{} merupakan sebuah type data
	var secret interface{} // bisa menampung data apa saja. biasa disebut konsep Dynamic Typing

	secret = "Team Secret"
	fmt.Println(secret)
	
	secret = []string{"Dayum", "Haikal", "You're", "Fine"}
	fmt.Println(secret)
	// slice ini gak bisa di looping karena type datanya interface, bukan slice. KECUALI
	// kita casting dulu ke data aslinya.
	fmt.Println(secret.([]string)[:2])
	for _, each := range secret.([]string){
		fmt.Println(each)
	}


	secret = 420.69
	fmt.Println(secret)

	fmt.Println("===============================================================")

	// map with string key, and interface value
	var everything = map[string]interface{}{
		"Nama": "Haikal",
		"Umur": 23,
		"Hobbies": []string{"Gaming", "Basketball", "Reading", "Coding"},
	}

	// fmt.Println(everything["Hobbies"])

	for key, value := range everything {
		if key == "Hobbies" {
			fmt.Printf("%s : \n", key)
			for j, hobby := range value.([]string) { // value Hoobies kan interface berisi []slice. dan kalo mau ambil/loop harus di casting dulu.
				fmt.Printf("%d. %s\n", j+1, hobby)
			}
		} else {
			fmt.Printf("%s : %v\n", key, value)
		}
		
	}
	
	
	
	fmt.Println()

	// Memakai any sebagai interface kosong.
	// Untuk GO versi 1.18
	// var segalanya map[string]any

	// segalanya = map[string]any{
	// 	"Name": "Haikal",
	// 	"Age": 23,
	// 	"gender": "Laki-laki",
	// }


}

// Contoh penggunaan interface pada kasus "SOAL"

// Soal 1
// Segitiga "SAMA SISI"
type segitigaSamaSisi struct {
	alas, tinggi int
}

func (sss segitigaSamaSisi) luas() int {
	return (sss.alas * sss.tinggi) / 2
}

func (sss segitigaSamaSisi) keliling() int {
	return sss.alas + sss.alas + sss.alas
}

// Persegi Panjang
type persegiPanjang struct {
	panjang, lebar int
}

func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

func (pp persegiPanjang) keliling() int {
	return 2 * (pp.panjang + pp.lebar)
}

// Tabung
type tabung struct {
	jariJari, tinggi float64
}

func (tab tabung) volume() float64 {
	return math.Pi * math.Pow(tab.jariJari, 2) * tab.tinggi
}

func (tab tabung) luasPermukaan() float64 {
	return 2 * math.Pi * tab.jariJari * (tab.jariJari + tab.tinggi)
}

// Balok
type balok struct {
	panjang, lebar, tinggi int
}

func (bal balok) volume() float64 {
	return float64(bal.panjang * bal.lebar * bal.tinggi)
}

func (bal balok) luasPermukaan() float64 {
	var lp = 2 * (bal.panjang * bal.lebar + bal.panjang * bal.tinggi + bal.lebar * bal.tinggi)
	return float64(lp)
}

// Interface
type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func soal_1() {
	var bangunDatar hitungBangunDatar
	var bangunRuang hitungBangunRuang

	fmt.Println("=============== Soal 1 ===============")

	bangunDatar = segitigaSamaSisi{15,13}
	fmt.Println(`Rumus Luas Segitiga sama sisi adalah : L = (alas * tinggi) / 2 "ATAU" L = (a² / 4) √3`)
	fmt.Printf("Luas Segitiga sama sisi dengan alas %v dan tinggi %v adalah : %v\n", bangunDatar.(segitigaSamaSisi).alas, bangunDatar.(segitigaSamaSisi).tinggi, bangunDatar.luas())
	fmt.Println("Karena Segitiga sama sisi maka sisi == alas dan rumus kelilingnya adalah =  alas + alas + alas")
	fmt.Printf("Keliling segitiga sama sisi dengan alas %v adalah : %v\n", bangunDatar.(segitigaSamaSisi).alas, bangunDatar.keliling())

	fmt.Println()

	bangunDatar = persegiPanjang{10, 6}
	fmt.Printf("Luas Persegi Panjang : %v\n", bangunDatar.luas())
	fmt.Printf("Keliling Persegi Panjang dengan panjang %v dan lebar %v adalah : %v\n", bangunDatar.(persegiPanjang).panjang, bangunDatar.(persegiPanjang).lebar, bangunDatar.keliling())

	fmt.Println()

	bangunRuang = tabung{7,10}
	fmt.Printf("Volume Tabung : %.2f\n", bangunRuang.volume())
	fmt.Printf("Luas Permukaan Tabung : %.2f\n", bangunRuang.luasPermukaan())	

	fmt.Println()

	bangunRuang = balok{10, 7, 5}
	fmt.Printf("Volume Balok : %.2f\n", bangunRuang.volume())
	fmt.Printf("Luas Permukaan Balok : %.2f\n", bangunRuang.luasPermukaan())	
	
	fmt.Println()
}

// Soal 2
type phone struct{
	name, brand string
	year int
	colors []string
}

func (p phone) printGadget() string {
	var output string
	
	output += fmt.Sprintln("Name \t:", p.name)
	output += fmt.Sprintln("Brand \t:", p.brand)
	output += fmt.Sprintln("Year \t:", p.year)
	output += fmt.Sprintln("Colors \t:", strings.Join(p.colors, ", "))

	return output
}

type gadget interface {
	printGadget() string
}

func soal_2(){
	// New interface 
	var smartphone gadget = phone{"Samsung Galaxy Note 20", "Samsung", 2020, []string{"Mystic Bronze", "Mystic White", "Mystic Black"}}

	fmt.Println("=============== Soal 2 ===============")

	// Print Interface
	fmt.Println(smartphone.printGadget())
}

// Soal 3
func luasPersegi(sisi int, cond bool) interface{} {
	var text interface{}
	var lPersegi = sisi * sisi

	if sisi != 0 && cond  {
		text = fmt.Sprintf("Luas persegi dengan sisi %v cm. Adalah : %v cm", sisi, lPersegi)
	} else if !cond {
		text = fmt.Sprintf("%v", sisi)
	} else if sisi == 0 && cond {
		text = fmt.Sprintf("Maaf anda belum menginput sisi dari persegi")
	} else {
		text = fmt.Sprintf("%v", nil)
	}

	return text
}

func soal_3(){
	fmt.Println("=============== Soal 3 ===============")

	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	fmt.Println()
}

// Soal 4
func soal_4(){
	var prefix interface{}= "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6,8}
	var kumpulanAngkaKedua interface{} = []int{12,14}

	fmt.Println("=============== Soal 4 ===============")

	var textTambah = func(angka interface{}) (string, int){
		var text string
		var sum int = 0
		for i, num := range angka.([]int) {
			if i != len(angka.([]int))-1{
				text += fmt.Sprintf("%v + ", num)
				sum += num
			} else {
				text += fmt.Sprintf("%v", num)
				sum += num
			}
		}

		return text, sum
	}

	var text1, sum1 = textTambah(kumpulanAngkaPertama)
	var text2, sum2 = textTambah(kumpulanAngkaKedua)
	fmt.Printf("%v%s + %s = %v", prefix , text1, text2, sum1+sum2)
}

type dan interface{
	haha() string
	hihi() int
	hoho()

}