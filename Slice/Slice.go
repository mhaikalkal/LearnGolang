package slice

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func SliceFunc() {
	// Slice ini mirip array, tapi bukan
	// pokoknya yang [] gada isinya itu slice

	var panggil = []string{"aku", "dia", "Kita", "mereka", "kalian", "sendiri"} // Slice
	var makanan = [...]string{"kolak", "wafer", "Tahu", "Mie"} // Array
	var minuman = [3]string{"Kopi", "susu", "teh"} // array

	fmt.Println(panggil)
	fmt.Println(makanan)
	fmt.Println(minuman)

	fmt.Println(panggil[1]) // pemanggilan masih sama kaya array

	// ambil data index:0 s/d sebelum 4 (3) dari slice fruit
	var newPanggil = panggil[0:3] 
	// fmt.Println(newPanggil)
	// output [aku dia kita]

	// panggil[a,b]
	// panggil[0:0] == data kosong, Karena tidak ada index sebelum 0
	// panggil [:] == ambil semua data
	// panggil [lastIndex:lastIndex] == data kosong
	// panggil[4:0] == error. b harus lebih kecil dari a
	// panggil[2:] == semua index dimulai dari index ke-2
	// panggil[:2] == semua index sampai dengan sebelum index ke-2

	// fungsi append (menambahkan)
	var addPanggil = append(newPanggil, "Semuanya")

	fmt.Println("====================================================")
	fmt.Println(newPanggil)
	fmt.Println(addPanggil)

	// panggil -> newPanggil -> addPanggil
	// Slice, kalau anaknya dirubah, parent/slice referensinya juga ikut berubah
	// contoh: yang diubah cuma addPanggil aja, slice reference yang lain ikut keubah juga
	addPanggil[0] = "Sayang"
	fmt.Println("====================================================")
	fmt.Println(panggil)
	fmt.Println(newPanggil)
	fmt.Println(addPanggil)


	// copy()
	fmt.Println("====================================================")
	destination := make([]string, 3) // bikin slice yang isinya 3
	source := []string{"one", "last", "time", "i", "need", "to", "be", "the", "one", "who", "takes", "you", "home"}

	n := copy(destination, source)

	fmt.Println(source)
	fmt.Println(destination)
	fmt.Println(n)
	// karena isi destination < source maka, cuma 3 elemen yang ke copy ke destination
	// kalau source < destination maka, semua element nya ke copy.

	for i:=0; i < len(destination); i++{
		fmt.Println(destination[i])
	}

	// ambil substring, menggunakan slice
	var kota = "sukabumi"
	// fmt.Println(len(kota))
	
	kota = strings.Title(kota[:len(kota)-1]) + strings.ToUpper(kota[len(kota)-1:])

	fmt.Println(kota)
	
}

func SliceWithIteration() {
	lirik := []string{"one", "last", "time", "i", "need", "to", "be", "the", "one", "who", "takes", "you", "home"}

	fmt.Println(lirik)
	fmt.Println("")

	for _, subLirik := range lirik {
		fmt.Printf("%s\n", subLirik)

	}

	fmt.Println("")

	rand.Seed(time.Now().UnixNano())
	max := len(lirik)
	min := 0
	randomized := rand.Intn((max - min) + min)
	newLirik := lirik[randomized]
	
	fmt.Println(randomized)
	fmt.Println(newLirik)

}

func RNGPercentBased() {
	var pool = []string{
		"legend-1",
		"rare-1",
		"rare-2",
		"rare-3",
		"normal-1",
		"normal-2",
		"normal-3",
		"normal-4",
		"normal-5",
		"normal-6",
	}
	
	var RNGRes = RNG(pool)	
	fmt.Println("You just got yourself a", RNGRes)

}

// Simple RNG, yang sebenernya enggak kaya gini.
func RNG(RNGPool []string) string{
	rand.Seed(time.Now().UnixNano())	
	rng := rand.Intn(100 - 1) + 1 // rand.Intn(max - min) + min

	// legend = 1%
	// rare = 30%
	// normal = 69%

	var res string
	if rng  == 1 {
		// fmt.Sprintf("%s", RNGPool[0])
		res = RNGPool[0]
	} else if rng > 1 && rng <= 32 {
		rare := RNGPool[rand.Intn(3 - 1) + 1]
		res = rare
	} else {
		norm := RNGPool[rand.Intn((len(RNGPool)-1) - 4) + 4]
		res = norm
	}

	return res

}