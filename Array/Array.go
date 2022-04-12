package array

import (
	"fmt"
)

func ArrayFunc() {
	// defining array
	var names [4]string
	names[0] = "Trafalgar"
	names[1] = "D"
	names[2] = "Law"
	names[3] = "BANZAIII"

	fmt.Println(names[0], names[1], names[2], names[3])

	// Pengisian array yang melebihi alokasi awal, akan error
	var abjad [4]string
	abjad[0] = "A"
	abjad[1] = "B"
	abjad[2] = "C"
	abjad[3] = "D"
	// abjad[4] = "E"  -- Baris kode ini menghasilkan error

	
	// Inisialisasi nilai awal
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t:", len(fruits)) // \t = tab
	fmt.Println("Isi semua element \t:", fruits)

	
	// horizontal
	var buah [4]string
	buah = [4]string{"Rambutan", "Melon", "Pisang", "Tomat"}

	fmt.Println("array buah : ", buah)

	// vertikal
	var number [4]int
	number = [4]int{
		3,
		5,
		7,
		1,
	}

	fmt.Println("array number : ", number)

	// inisialisasi tanpa panjang array/jumlah array
	var numbers = [...]int{5,6,2,36,24,7,35,2,88,350,3,112}

	// akan otomatis keisi jumlah array elemennya
	fmt.Println("Jumlah element \t\t:", len(numbers))
	fmt.Println("Isi semua element \t:", numbers)

	// memakai make
	var sayur = make([]string, 3) // buat array of string yang isinya 3

	sayur[0] = "Carrot"
	sayur[1] = "Bok Choy"
	sayur[2] = "Egg Plant"

	fmt.Println(sayur)

	// string with array func test
	var furnitures = [...]string{"kursi", "meja", "lemari", "kabinet", "bangku"}
	fmt.Println(string(furnitures[0][1]))

}

func ArrayFuncWithIteration() {
	var phones = [...]string{"samsung", "iphone", "nokia", "siemens", "xiaomi"}

	for i:= 0; i < len(phones); i++{
		fmt.Println("Aku punya HP : ", phones[i])

	}

	// pake range
	for j, phone := range phones{
		fmt.Printf("HP ke-%d : %s\n", j+1, phone);
	}

	// Kalau mau isi array-nya aja pake var pengangguran _ biar ga ada variable yang gak kepake
	for _, ph := range phones{
		fmt.Printf("Aku punya handphone merk %s\n", ph);
	}

	// for _, phone := range phones --- kalau mau isi array nya aja
	// for i, _ := range phones  --- kalau mau index array nya aja

	var furnitures = [...]string{"kursi", "meja", "lemari", "kabinet", "bangku"}

	for y, furniture := range furnitures{
		fmt.Printf("furniture ke-%d memiliki %d huruf :\n", y+1, len(furniture))

		for x, abjad := range furniture{
			fmt.Printf("huruf ke-%d : %s\n", x+1,  string(abjad))
		}

		fmt.Println("")
		
	}

	
}

func ArrayMultidomension() {
	// 2 array, yang isinya 4
	num := [2][4]int{{1,2,3}, {6,7,8}} // karena diisi 3, index ke-3 nya jadi 0

	fmt.Println(len(num))
	fmt.Println(num)
	fmt.Println("=================================================================================================================")
	
	/* 
		Kita punya 2 dus Kopi.
		Per 1 dus ada 4 bungkus.
		Per 1 bungkus ada 3 sachets.
	*/

	nums := 
	[2][4][3]int{
		{
			{11, 99, 88},
			{22, 66, 77},
			{33, 100},
		},
		{
			{5, 1, 2},
			{6, 3},
			{7},
			{12412, 7782, 4089},
		},
		
	}

	// fmt.Println(nums)
	fmt.Println("Saya punya", len(nums), "Dus.")
	fmt.Println("=================================================================================================================")

	for x, dus := range nums{
		fmt.Printf("Dus ke-%d isi ada : ", x+1)
		fmt.Println(len(dus), "bungkus")
		for y, bungkus := range dus{
			fmt.Printf("\tBungkus ke-%d isi ada : ", y+1)
			fmt.Println(len(bungkus), "sachets")
			for z, pcs := range bungkus{
				fmt.Printf("\t\tSachet ke-%d isi ada : ", z+1)
				fmt.Println(pcs, "buah")
			}
		}
	}

	fmt.Println("=================================================================================================================")
	fmt.Println("Penjumlahan Isi Array")
	newNum := []int{2,4,5}

	result := 0
	for o := 0; o < len(newNum); o++{
		result += newNum[o]
	}
	
	fmt.Println(result)
}