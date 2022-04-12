package variable

import (
	"fmt"
	"strconv"
	"strings"
)

func VarFunc() {
	// ada 3 cara declare variable
	// 1. Berguna jika mau manggil beda Scope
	var i int
	i = 33

	// 2. Lengkap, biar compiler ga bingung
	var j float64 = 29

	// 3. Compiler yang nge decide type nya apa
	k := 45.

	// multi declare variable
	var (
		firstName string = "blabla"
		lastName string = "hahaha"
		age int = 23
	)

	var first, second, third string
	first, second, third = "1", "2", "3"

	var fourth, fifth, sixth string = "empat", "lima", "enam"

	seventh, eighth, ninth := "tujuh", "delapan", "sembilan"

	fmt.Println(first, second, third, fourth, fifth, sixth, seventh, eighth, ninth)

	// _ (underscore) adalah reserve variable untuk nilai yang tidak dipakai
	// cukup pakai = buat assign
	// kalo multiple pake :=
	_ = "Haikal belajar golang"
	_ = "Golang mantap cuy"
	tenth, _ := "seribu", 23
	fmt.Println(tenth)

	var umur string = strconv.Itoa(age)
	fmt.Println("My name is " +firstName +" " +lastName +" Umur saya " +umur + " tahun")
	
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)

}

func TugasJCC_2() {
	// Soal 1
	var kata1, kata2, kata3, kata4, kata5 string = "Jabar", "Coding", "Camp", "Golang", "2022"

	fmt.Println(kata1, kata2, kata3, kata4, kata5);


	// Soal 2
	halo := "Halo Dunia"
	var golang = strings.Replace(halo, "Dunia", "Golang", 1); 
	// kalau misal halo nya "halo Dunia Dunia" dan n=1 output: "halo Golang Dunia", kalo n=2 output: "halo Golang Golang"

	fmt.Println(golang)
	

	// Soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	newKataKedua := strings.Title(kataKedua)
	newKataKetiga := strings.Replace(kataKetiga, "r", "R", 1)
	newKataKeempat := strings.ToUpper(kataKeempat)

	fmt.Println(kataPertama, newKataKedua, newKataKetiga, newKataKeempat)


	// Soal 4
	var angkaPertama= "8"
	var angkaKedua= "5"
	var angkaKetiga= "6"
	var angkaKeempat = "7"

	firstNumber,_ := strconv.Atoi(angkaPertama)
	secondNumber,_ := strconv.Atoi(angkaKedua)
	thirdNumber,_ := strconv.Atoi(angkaKetiga)
	lastNumber,_ := strconv.Atoi(angkaKeempat)

	grandTotal := firstNumber + secondNumber + thirdNumber + lastNumber
	fmt.Println("Total dari 4 variable diatas adalah:", grandTotal)


	// Soal 5
	kalimat := "halo halo bandung"
	angka := 2022

	firstKalimat := strings.Replace(kalimat, "alo", "i", 2)
	newKalimat := strings.Replace(firstKalimat, "h", "H", 2)
	fmt.Printf(`"%s" - %d`, newKalimat, angka)


	namaku := "Haikal"
	umur := 23

	fmt.Println(namaku, umur)


	var sentence = "hahaikal ."
	var dot = sentence[len(sentence)-1:]
	fmt.Println(dot)

	Buah()
}

func Buah(){
	fmt.Println("=============== SOAL 3 ===============")
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("john", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	fmt.Println()
}

func buahFavorit(name string, fruits ...string) string{
	var fruitsToString string
	// fmt.Println("ada =", len(fruits))
	for x, fruit := range fruits {
		if x == len(fruits)-1{
			fruitsToString += fmt.Sprintf(`"%s" `, fruit)
		} else {
			fruitsToString += fmt.Sprintf(`"%s", `, fruit)
		}
	}

	var sentence = fmt.Sprintf(`halo nama saya %s dan buah favorit saya adalah %s`, name, fruitsToString)

	return sentence
}