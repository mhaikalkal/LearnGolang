package pointer

import (
	"fmt"
	"math"
)

func PointerTest() {
	/* 
		1. * (Pointer) isinya Address / value yg diubah jadi address menggunakan &
		* bisa membaca address dan menerjemahkannya ke value asli

		Misal: var address berisi `address`.
		var address = 0xc000014070
		fmt.Println(*address)
		output = 9
	*/

	var numberA int = 4
	var numberB *int = &numberA // var numberB merupakan pointer yang mereferensi numberC
	
	fmt.Println(numberA)
	fmt.Println(numberB)

	fmt.Println(&numberA) // membaca type int menjadi address
	fmt.Println(*numberB) // membaca address menjadi value biasanya

	/*
		2. & tanda ini mereferensi.
		& bisa membaca value dan menerjemahkannya ke address.

		Misal: 
		var numberC = 9
		fmt.Println(&numberC)
		output = 0xc0000ae0a0     atau     address lain.

	*/

	var numberC int = 9
	var numberD *int = &numberC

	fmt.Println("========================================")
	fmt.Println("NumberC Value\t:", numberC)
	fmt.Println("NumberC Address\t:", &numberC)
	
	fmt.Println("NumberD Value\t:", *numberD)
	fmt.Println("NumberD Address\t:", numberD)

	fmt.Println("========================================")

	numberC = 10
	fmt.Println("NumberC Value\t:", numberC)
	fmt.Println("NumberC Address\t:", &numberC)
	
	fmt.Println("NumberD Value\t:", *numberD)
	fmt.Println("NumberD Address\t:", numberD)

}


// Soal 1
func soal1(){
	fmt.Println("=============== Soal 1 ===============")
	var luasLingkaran float64 
	var kelilingLingkaran float64

	/// jari-jari
	var jari float64 = 7

	// Masukin reference ke func() diterima sebagai pointer
	var Lingkaran = func(jari float64, luasLingkar, kelilingLingkar *float64) (float64, float64){
		luas := luasLingkar
		// luas mereferensi si luasLingkar
		*luas = math.Pi * math.Pow(jari, 2)
		// *luas artinya Value luas = math.Pi * math.Pow(jari, 2)
		// kalau kita tulis luas error, karena si luas skrg hanya address.
		
		keliling := kelilingLingkar
		*keliling = 2 * math.Pi * jari

		// karena return nya pake *
		// return *luas, *keliling === dibaca return value luas, keliling
		// type return di sebelah func nya jadi biasa (float64, float64)
		return *luas, *keliling

		// kalau return nya tanpa *
		// return luas, keliling === dibaca return address luas, address keliling
		// type return di sebelah func nya jadi (*float64, *float64) // di convert lagi jadi value
	}

	Lingkaran(jari, &luasLingkaran, &kelilingLingkaran)

	fmt.Printf("Luas Lingkaran dengan jari-jari %v adalah : %.2f\n", jari, luasLingkaran)
	fmt.Printf("Keliling Lingkaran dengan jari-jari %v adalah : %.2f\n", jari, kelilingLingkaran)

	fmt.Println()

}

// Soal 2
func soal2(){
	fmt.Println("=============== Soal 2 ===============")
	
	var sentence string 
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	fmt.Println()

}

func introduce(sentence *string, name, sex, job, age string) *string{
	kalimat := sentence
	if sex == "laki-laki" {
		*kalimat = fmt.Sprintf("Pak %s adalah seorang %s yang berusia %s tahun", name, job, age)
	} else {
		*kalimat = fmt.Sprintf("Bu %s adalah seorang %s yang berusia %s tahun", name, job, age)
	}

	return kalimat
}

// Soal 3
func soal3(){
	fmt.Println("=============== Soal 3 ===============")

	var buah = []string{}
	var listBuahBaru = [...]string{"Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat"}

	for _, buahbaru := range listBuahBaru{
		TambahBuah(&buah, buahbaru)
	}
	// TambahBuah(&buah, "buahbaru")
	for x, buahAdded := range buah {
		fmt.Printf("%v. %s\n", x+1, buahAdded)
	}

	fmt.Println()
}

func TambahBuah(fruits *[]string, listBuahBaru string) []string {
	var buahNew *[]string = fruits // buahNew mereferensi address fruit/buah
	*buahNew = append(*buahNew, listBuahBaru) // append(destination, data), dibaca vvv
	// value buahNew diisikan (value buahNew, listBuahBaru)

	// karena disini pake * maka type nya tanpa *
	return *buahNew
}

// Soal 4
func soal4() {
	fmt.Println("=============== Soal 4 ===============")

	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	// fmt.Println(dataFilm)

	for i, movies := range dataFilm{
		fmt.Printf("%v. ", i+1)
		for j, movie := range movies {
			fmt.Printf("%s : %s\n", j, movie)
		}
		
		fmt.Println()
	}
}

func tambahDataFilm(title, duration, genre, year string, dataFilm *[]map[string]string) *[]map[string]string {
	// dataMovies mereference slice of map dataFilm
	var dataMovies *[]map[string]string = dataFilm

	// karena didalam slice isinya map, maka kita declare map listMovie, untuk mengisi slice.
	// step ini normal tanpa pointer.
	var listMovie = map[string]string{"title": title, "duration": duration, "genre": genre, "year": year }

	// value dataMovies diisi append(value dataMovies, map listMovie)
	*dataMovies = append(*dataMovies, listMovie)

	// karena tidak memakai * maka type return nya menjadi *map[string]string
	return dataMovies
}