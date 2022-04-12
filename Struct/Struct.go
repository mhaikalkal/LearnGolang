package main

import (
	"fmt"
)

/*
	Konsep struct di golang mirip dengan konsep class pada OOP, meski sebenarnya berbeda.
	Di sini penulis menggunakan konsep OOP sebagai analogi, dengan tujuan untuk mempermudah dalam mencerna isi chapter ini.

	Struct ini mirip type data map.
*/

// Declare Struct
func main(){
	// How to Declare and assign a Struct
	studentStruct()
	
	// Embedded Struct
	teacherStruct()
	
	// Embedded Struct with same property(s)
	mobsStruct()

	// Anonymous Struct
	anonymStruct()

	// Struct + Slice
	sliceAnonymStruct()

	// Nested Struct
	nestedStructManusia()

	// Alias Struct
	aliasStruct()
}

type student struct {
	name string
	grade int
}

func studentStruct(){
	// Inisialisasi Struct
	// Cara 1
	var st1 = student{}
	st1.name = "Light Yagami"
	st1.grade = 98
	
	// Cara 2
	var st2 = student{"Misa Misa", 65}

	// Cara 3
	var st3 = student{name:"Ryuk"}

	// Tidak berurutan
	var st4 = student{grade: 44, name: "Sakurai"}

	fmt.Println(st1)
	fmt.Println("Name :", st1.name)
	fmt.Println("Grade:", st1.grade)
	fmt.Println()

	fmt.Println(st2)
	fmt.Println("Name :", st2.name)
	fmt.Println("Grade:", st2.grade)
	fmt.Println()
	
	fmt.Println(st3)
	fmt.Println("Name :", st3.name)
	fmt.Println("Grade:", st3.grade)
	fmt.Println()

	fmt.Println(st4)
	fmt.Println("Name :", st4.name)
	fmt.Println("Grade:", st4.grade)

}


// Embedded Struct
type person struct {
	name string
	age int
}

type teacher struct{
	person
	nip string
}

func teacherStruct() {
	var teach1 = teacher{}
	teach1.name = "Koro Sensei"
	teach1.age = 99
	teach1.nip = "112233"

	fmt.Println("======================================================")

	fmt.Println("Nama \t:", teach1.name) // meskipun yang diisi cuma teach1.name, teach1.person.name akan mengikuti
	fmt.Println("Nama (Person) :", teach1.person.name)

	fmt.Println("Umur \t:", teach1.age) // dipanggil pun bisa tanpa pakai .person, karena struct person menjadi bagian dalam teacher juga
	fmt.Println("Umur (Person) :", teach1.person.age)

	fmt.Println("NIP \t:", teach1.nip)

	fmt.Println()
}

// Embedded Struct dengan nama property yang sama
type monsters struct {
	name string
	age int // sama mobs
}

type mobs struct {
	monsters
	age int // sama monsters
	element string
}

func mobsStruct() {
	var mob1 = mobs{}

	mob1.name = "Darion"
	mob1.age = 27
	mob1.monsters.age = 33 // Jika ini tidak diisi, maka monsters.age = 0 (sesuai type data int)
	mob1.element = "Light"

	fmt.Println("======================================================")

	fmt.Println("Mob name \t :", mob1.name)
	fmt.Println("Mob age \t :", mob1.age) // meski mobs.age diisi, monsters,age tidak akan mengikuti, meskipun nama property-nnya sama
	fmt.Println("Mob Monsters age :", mob1.monsters.age)
	fmt.Println("Mob element \t :", mob1.element)

	fmt.Println()
}

// Anonymous Struct
type orang struct {
	name string
	age int
}

// untuk var anonym
type spesifikasi struct {
	kapasitas int
	biaya int
}

// Cara ini bagus buat kasus yang bikin struct cuma 1x 
func anonymStruct(){
	// langsung bikin var isinya struct
	var anonim = struct {
		orang
		job string
	}{} // Cara 1. Inisialisasi seperti biasa

	anonim.name = "Anonim"
	anonim.age = 17
	anonim.job = "Assasins"

	fmt.Println("======================================================")

	fmt.Println(anonim.name)
	fmt.Println(anonim.age)
	fmt.Println(anonim.orang.age)
	fmt.Println(anonim.job)

	fmt.Println()

	var org1 = struct {
		orang
		hobby string
	}{
		orang: orang{"Haikal", 22},
		hobby: "Basketball",
	} // Cara 2. Inisialisasi langsung
	
	fmt.Println("======================================================")

	fmt.Println(org1.name)
	fmt.Println(org1.age)
	fmt.Println(org1.orang.age)
	fmt.Println(org1.hobby)
	
	fmt.Println()

	// Cara 3
	// Menggunakan var
	var rumah struct{
		alamat string
		nomer int
		spesifikasi
	}

	rumah.alamat = "Selabintana"
	rumah.nomer = 40
	rumah.kapasitas = 8
	rumah.biaya = 5000000000

	fmt.Println("======================================================")

	fmt.Println(rumah.alamat)
	fmt.Println(rumah.nomer)
	fmt.Println(rumah.kapasitas)
	fmt.Println(rumah.biaya)	

	fmt.Println()
}

// Anonym Struct with Slice
type siswa struct {
    name string
    age int
}

func sliceAnonymStruct(){
	// Cara 1
	/// kalau struct-nya udah ada
	// dan gak ada embedded. contoh embedded slice struct ada dibawah ini
	var allSiswa = []siswa{
		{name: "Haikal", age: 12},
		{name: "Fay Nabila", age: 11},
		{name: "Gia", age: 11},
		{name: "Kira", age: 13},
	}
	
	fmt.Println("======================================================")	
	for _, student := range allSiswa {
		fmt.Println(student.name, "age is", student.age)
	}

	fmt.Println()

	// Cara 2
	var allMahasiswa = []struct{
		siswa
		prodi string
		semester int
	}{
		// karena embedded siswa ke struct allMahasiswa, maka cara ngisinya panggil dulu struct siswa-nya. Baru isi kaya biasa
		{siswa: siswa{"Haikal",22}, prodi: "Teknik Informatika", semester: 5},
		{siswa: siswa{"Sandi",25}, prodi: "Teknik Sipil", semester: 9},
		{siswa: siswa{"Alam",20}, prodi: "Teknik Industri", semester: 1},
	}

	fmt.Println("======================================================")
	for i, mahasiswa := range allMahasiswa {
		fmt.Printf("%v. %s. Berumur %d sedang menjalankan kuliah pada semester %d dengan jurusan %s.\n", i+1, mahasiswa.siswa.name, mahasiswa.siswa.age, mahasiswa.semester, mahasiswa.prodi)	
	}

	fmt.Println()
}

// Nested Struct
// nested struct merupakan anonymous struct yang embednya dideclarasikan didalam struct. pusing ajg
// intinya anonym struct di embed, dan ada/bisa declare struct baru juga
type identity struct { name string; age int }

func nestedStructManusia(){
	// Cara 1
	var manusia struct{
		identity 
		alamat []string `tag` // ini tag, buat nanti JSON
		activity struct{ // langsung bikin activity type data struct, kaya declare `nama string`
			job, hobby string
		}
	}

	// Cara 2
	// var manusia = struct{
	// 	identity
	// 	alamat []string
	// 	activity struct{ job, hobby string }
	// }{} // Diisi disini langsung juga bisa

	manusia.identity.name = "Muhammad Haikal"
	manusia.identity.age = 23
	manusia.alamat = []string{"Jl. Selabintana no. 40", "Kota Sukabumi"}
	manusia.activity.hobby = "Basketball"
	manusia.activity.job = "none"

	fmt.Println("======================================================")
	
	fmt.Println("Nama :", manusia.identity.name)
	fmt.Println("Umur :", manusia.identity.age)

	for i, tempat := range manusia.alamat{
		if i == 0 {
			fmt.Printf("Alamat : %s\n", tempat)
		} else {
			fmt.Printf("Domisili : %s\n", tempat)
		}
	}

	fmt.Println("Hobby :", manusia.activity.hobby)
	fmt.Println("Pekerjaan :", manusia.activity.job)

	fmt.Println()

}

// Alias
type handphone struct{
	name, brand string

}

type HP = handphone
// Dengan melakukan ini, HP dan handphone memiliki struktur yang sama

func aliasStruct(){

	var hp1 = handphone{ name:"Samsung Galaxy Z Fold", brand:"Samsung" }
	var hp2 = HP{ name:"POCO M3", brand:"XIAOMI" }

	fmt.Println("======================================================")

	fmt.Println(hp1)
	fmt.Println(hp2)
	
	// var handheld = HP{parameter}
	// HP parameter dimasukkan ke struct handphone.
	handheld := HP{"Iphone 7", "Apple"}
	fmt.Println(handphone(handheld))

	smartphone := handphone{"Infinix Hot 10S Pro", "Infinix"}
	fmt.Println(HP(smartphone))

	fmt.Println()

	// Alias juga bisa dipakai sebagai berikut. dan masih banyak lagi
	type Number = int
	var umur Number = 23

	fmt.Printf("%T, %v", umur, umur)
}

