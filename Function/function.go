package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	Greet()
	Kubus()
	Circle()
	Square()
	Avg()
	// Closure
	Closure()
	Film()
	// Closure IIFE
	FilterNumber()
	// Func yange mereturn func
	FilterSiswa()

	// CALLBACK
	Mahasiswa()
}

// tanpa mengembalikan type data
func Greet(){
	nama := []string{"Muhammad", "Haikal"}
	greet := "hello, my name is "
	ShowGreetings(greet, nama)

	firstName, lastName := introduction("John", "Doe")
	fmt.Println(firstName, lastName)

	// menghiraukan return. dibuang return nya
	firstName2, _ := introduction("John", "Doe")
	fmt.Println(firstName2)
	fmt.Println("======================================================")

}

// tanpa return, si function yg nge print nya
func ShowGreetings(greet string, arrNama []string) {
	var name = strings.Join(arrNama, " ")
	fmt.Println(greet, name)
}

func introduction(firstName string, lastName string) (string, string) {
	introFirstName := "Hello My First Name Is " + firstName
	introLastName := "Hello My Last Name Is " + lastName
	return introFirstName, introLastName
}


// Volume kubus. return 1 data.
func Kubus(){
	var sisi = 5
	var VolumeKubus = RumusLuasKubus(sisi)
	fmt.Println(VolumeKubus)
	fmt.Println("======================================================")
}

// non-defined single return
func RumusLuasKubus(sisi int) int {
	var value = sisi*sisi*sisi
	return value
}


// Lingkaran. return multiple data
func Circle() {
	var diameter float64 = 8
	var luasLingkaran, kelLingkaran = RumusLingkaran(diameter)
	fmt.Printf("Luas lingkaran yang berdiameter %.2f = %.2f\n", diameter, luasLingkaran)
	fmt.Printf("Keliling lingkaran yang berdiameter %.2f = %.2f\n", diameter, kelLingkaran)
	fmt.Println("======================================================")
}

// non-defined multiple return
func RumusLingkaran(d float64) (float64, float64) {
	var luasLingkaran = math.Pi * math.Pow(d/2, 2) // pi * r^2 / pi * (diameter/2)^2
	var kelLingkaran = math.Pi * d

	return luasLingkaran, kelLingkaran
}


// persegi, Predefined return value
func Square(){
	var panjang = 5
	var kelPersegi, luasPersegi = RumusPersegi(panjang)
	fmt.Printf("Luas persegi yang memiliki panjang %d = %d\n", panjang, luasPersegi)
	fmt.Printf("Keliling lingkaran yang memiliki panjang %d = %d\n", panjang, kelPersegi)
	fmt.Println("======================================================")
}

// Predefined multiple return
func RumusPersegi(p int) (luasPersegi int, kelPersegi int) {
	luasPersegi = p * p
	kelPersegi = 2 * (p + p)

	// return kosong, artinya menghentikan function
	return
}


// fungsi variadic
func Avg(){
	var avg = Calculate(10, 20, 30, 40, 50)
	var msg = fmt.Sprintf("Rata-rata = %.2f", avg)
	fmt.Println(msg)
	fmt.Println("======================================================")

}

// Variadic Function
func Calculate(numbers ...int) float64{
	var total int = 0
	for _, number := range numbers{
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

// Function CLOSURE, func yang dibungkus kedalam variable
func Closure() {
	var getMinMax = func(n []int) (int, int) {
        var min, max int
        for i, e := range n {
            switch {
            case i == 0: // index cuma 1, ke-0
                max, min = e, e // maka min dan max sama
            case e > max:
                max = e
            case e < min:
                min = e
            }
        }
        return min, max
    }

    var numbers = []int{2, 3, 4, 3, 4, 2, 3}
    var min, max = getMinMax(numbers)
    fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)
	fmt.Println("======================================================")
}

func Film(){
	var dataFilm = []map[string]string{}
	var tambahDataFilm = func(title, jam, genre, tahun string) []map[string]string {
		var films = map[string]string{"title": title, "jam": jam, "genre": genre, "tahun": tahun } // bikin isi slice-nya
		dataFilm = append(dataFilm, films) // append isi slice ke slice ORI-nya

		return dataFilm
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}

	fmt.Println("======================================================")
}


// IIFE = Imediatelly Invoke Function Expression. (Langsung di execute)
// biasanya dipakai untuk fungsi yang ke trigger cuma sesekali
func FilterNumber(){
	var numbers = []int{10, 9, 8, 7, 6, 200, 4, 3, 2, 1, 1000}
	// atau bisa juga biar si fungsi ga nerima filter diawal
	var f = 7

	var filtering = func(filter int) []int { // parameter filternya nanti dibawah
		var filtered = []int{}

		for _, num := range numbers {
			if num < filter { // jika number < filter/angka penyaring,
				continue // maka lewat/ kecualikan. jadi nomer dibawah angka filter ilang. ga di print
			}

			filtered = append(filtered, num) // masukin angka > filter ke slice filtered

		}

		return filtered

	}(f) // (5) parameter filter-nya, bisa variable, atau manual langsung nomer

	fmt.Println("Angka sebelum di filter \t:", numbers)
	fmt.Println("Filter \t\t\t\t: <", f)
	fmt.Println("Angka yang sudah di filter \t:", filtering)
	fmt.Println("======================================================")

}

// Function sebagai return
func FilterSiswa(){
	var names = []string{
		"Muhammad Haikal",
		"Muhammad Alfiansyah P",
		"Monica Angelina",
		"Ayu Dewi Putri",
		"I Made Ketut D Radja",
		"Muhammad Iqbal Pratama",
		"Fajar Putra Maharaja",
		"Veronica Kusumaatmadja",
		"Syifa Rahayu",
	}
	var NameContain = "Muhammad"
	var NameRows, NameFiltered = FilterASliceString(names, NameContain)

	fmt.Printf("Jumlah Siswa : %d\n", len(names))
	fmt.Println("Data Seluruh Siswa : ")
	for i, siswa := range names {
		fmt.Printf("%d. %s\n", i+1, siswa)
	}

	fmt.Println()
	fmt.Println("Siswa yang bernama :", NameContain)
	fmt.Println("Jumlah Siswa :", NameRows)
	var SiswaFiltered = fmt.Sprintf("Hasil : %v", NameFiltered()) // Panggil Func yang di return FilterASliceString
	fmt.Println(SiswaFiltered)
}

func FilterASliceString(AllSiswa []string, NamaSiswa string) (int, func() string){ // return int, dan fungsi yang nge return string
	var siswa = []string{}
	for _, name := range AllSiswa{
		if strings.Contains(name, NamaSiswa){ // jika slice siswa, mempunyai string NamaSiswa="Muhammad"
			siswa = append(siswa, name) // maka append siswa itu ke slice siswa
		}
	}
	
	// fungsi ini nge return,
	// (int, func() string)
	return len(siswa), func() string { // tadinya mau nge return func() []string. tapi jadi kurang bagus soalnya belum di olah slice nya
		// data siswa yang udah ke filter sama strings.Contains()
		var allNameFiltered string
		for x, nf := range siswa{
			if x == 0{ // kalau x == 0, maka kasih newline di awal kalimat, biar rapih.
				allNameFiltered += fmt.Sprintf("\n%d. %s\n", x+1, nf)
			} else {
				allNameFiltered += fmt.Sprintf("%d. %s\n", x+1, nf)
			}
		}
		return allNameFiltered
	}

}


// Fungsi sebagai parameter.
func Mahasiswa() {
	fmt.Println("======================================================")

	var allmahasiswa = []string{
		"Muhammad Haikal",
		"Muhammad Alfiansyah P",
		"Monica Angelina",
		"Ayu Dewi Putri",
		"I Made Ketut D Radja",
		"Muhammad Iqbal Pratama",
		"Fajar Putra Maharaja",
		"Veronica Kusumaatmadja",
		"Syifa Rahayu",
		"Dermawan Kitagawa",
		"Jocco Ferdiawan",
		"Gerrald Prindafan Batubara",
	}

	var filterHuruf = "o"
	var namaContainsO = filterMahasiswa(allmahasiswa, func(mahasiswa string) bool{ // Masukin slice allMahasiswa, dan callback yg menerima nama mahasiswa dan nge return true biar bisa diolah.
		return strings.Contains(mahasiswa, filterHuruf) // ini nge return BOOL.
		// Setiap Looping. Dikirim lagi ke func filterMahasiswa.
	})

	fmt.Printf("Jumlah Mahasiswa yang namanya mengandung huruf %s berjumlah : %d Mahasiswa\n", filterHuruf, len(namaContainsO))
	for i, FilteredMahasiswa := range namaContainsO {
		fmt.Printf("%v. %s\n", i+1, FilteredMahasiswa)
	}


}

func filterMahasiswa(allmahasiswa []string, callback func(mahasiswa string) bool ) []string { 
	var mahasiswaFiltered = []string{}
	for _, mahasiswa := range allmahasiswa {
		// pemanggilan callbacknya
		if filtered := callback(mahasiswa); filtered {  // filtered ini nerima return dari callback.
			// jika nama mahasiswa Contains filterHuruf (true), maka mahasiswa tersebut di append ke mahasiswaFiltered,
			mahasiswaFiltered = append(mahasiswaFiltered, mahasiswa)
		}
	}
	return mahasiswaFiltered

}