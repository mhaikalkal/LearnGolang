package main

import (
	"fmt"
	"reflect"
)

func main() {
	// reflectFund()
	// reflectStruct()
	InfoMethodVarStruct()
}

func reflectFund() {
	/*
		Reflection adalah teknik untuk inspeksi variabel, mengambil informasi dari variabel tersebut atau bahkan memanipulasinya.
		Cakupan informasi yang bisa didapatkan lewat reflection sangat luas, seperti melihat struktur variabel, tipe, nilai pointer, dan banyak lagi.

		Dari package reflect, 2 fungsi yang paling penting / paling sering digunakan adalah : 
		- reflect.ValueOf() akan mengembalikan objek dalam tipe reflect.Value, yang berisikan informasi yang berhubungan dengan nilai pada variabel yang dicari
		- reflect.TypeOf() mengembalikan objek dalam tipe reflect.Type. Objek tersebut berisikan informasi yang berhubungan dengan tipe data variabel yang dicari
	*/

	var number = 23
	// ambil detail data number
	var dataNumber = reflect.ValueOf(number) 

	// return type data `var dataNumber` : bisa pake dataNumber.Type() / dataNumebr.Type()
	fmt.Println("tipe variabel :", dataNumber.Kind())

	// dataNumber.Kind() return type data. Maka dija dataNumber.Sejenis(type data) == reflect.Int (int).
	if dataNumber.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", dataNumber.Int())
	}

	fmt.Println()

	// Mengakses nilai variable menggunakan Interface.
	var numb = 21
	// variable yang ingin menggunakan .Interface/Type/Kind harus udah di reflect.ValueOf() dulu
	var dataNumb = reflect.ValueOf(numb)

	fmt.Println("tipe  variabel :", dataNumb.Type())
	fmt.Println("Numb nilai variabel :", dataNumb.Interface())
	
	fmt.Println()
	
	// Fungsi Interface() mengembalikan nilai interface kosong atau interface{}.
	// Nilai aslinya sendiri bisa diakses dengan meng-casting interface kosong tersebut.
	var numbClone = dataNumb.Interface().(int)
	fmt.Println("Numb Clone nilai variabel :", numbClone)

	fmt.Println("=============================================================")
}

func reflectStruct() {
	var stud student
	// var stude student
	// var studen student

	stud = student{"Haikal", 23}
	// stude = student{"Venom", 88}
	var studen = &student{"Watson", 21}

	stud.getPropertyInfo()
	// stude.getPropertyInfo()
	studen.getPropertyInfoPtr()

	// coba tanpa method
	var chu = student{"Maha", 16}

	fmt.Println("===== Data Struct =====")
	z := reflect.ValueOf(&chu).Elem() // untuk mengambil pointer menggunakan elem
	fmt.Println(z) // output: {value Name, value age}

	fmt.Println(z.Type())
	// output = main.student
	// Artinya: TypeOf(n) == ValueOf(n).Type()
	// maka untuk ambil field struct. main.student.Field(index) kan. nah jadi bisa pake salah satu diatas.
	
	for x := 0; x < z.NumField(); x++ {
		varName := z.Type().Field(x).Name
		varValue := z.Field(x).Interface()
		fmt.Printf("%v : %v\n", varName, varValue)
	}

	fmt.Println("===== End Data Struct =====")
}

// Struct reflect
type student struct{
	Name string
	Age int
}

func (s student) getPropertyInfo() {
	var sDataValueOf = reflect.ValueOf(s)
	var sDataTypeOf = reflect.TypeOf(s)

	fmt.Println("===== Data Struct =====")
	
	fmt.Println("Data (s) : ", s) 
	fmt.Println("reflect.ValueOf(s) :", sDataValueOf)
	fmt.Println("reflect.TypeOf(s) :", sDataTypeOf)

	// Karena sekarang sudah bisa pake len (NumField) maka bisa di looping struct nya
	// Field(index).Name = ambil Field Name tiap-tiap index.
	// Field(index).Interface = ambil value tiap index, kenapa interface? karena isi value tiap field bisa beda-beda.
	for j := 0; j < sDataTypeOf.NumField(); j++ {
		fmt.Println(sDataTypeOf.Field(j).Name, ":", sDataValueOf.Field(j).Interface())
	}
	// Kenapa pas field name pake yg TypeOf ? karena dia ambil semua detail type data struct-nya.
	// Kenapa pas value pake yg ValueOf ? karena dia ambil semua value dari tiap index-indexnya

	fmt.Println("===== End Data Struct =====")

	fmt.Println()
}

func (s *student)getPropertyInfoPtr() {
	// Kalau gini (tanpa Pointer), kita ubah dulu jadi elemen.
	var reflectValue = reflect.ValueOf(s)

	fmt.Println("===== Data Struct =====")

	if reflectValue.Kind() == reflect.Ptr {
		fmt.Println("reflectValue.Kind() : ", reflectValue.Kind()) // Output: Ptr
		fmt.Println("reflectValue.Type() : ", reflectValue.Type()) // Output: *main.student

		reflectValue = reflectValue.Elem()

		fmt.Println(reflectValue) // Output: {value Name, value Age}
	}

	var reflectType = reflectValue.Type()

	fmt.Println("reflectValue.Type() after Elem() : ", reflectValue.Type()) // Output: main.student

	for i := 0; i < reflectValue.NumField(); i++{
		fmt.Println(reflectType.Field(i).Name, ":", reflectValue.Field(i))
		// ambil main.student.field(index) isinya : value.field(index)
	}

	fmt.Println("===== End Data Struct =====")

	fmt.Println()
}

// Method
func (s *student) SetStructStudent(name string, age int) {
	s.Name = name
	s.Age = age + 1000
}

func InfoMethodVarStruct(){
    var s1 = &student{Name: "john wick", Age: 2}
    fmt.Println("nama :", s1.Name)

    var reflectValue = reflect.ValueOf(s1)
	fmt.Println(reflectValue) // Output: &{john wick 2}

    var methodS1 = reflectValue.MethodByName("SetStructStudent") // address
	fmt.Println(&methodS1)	// Output: <func(string) value>

	var z = "Bebeb"

	// methodS1 == SetStructStudent(name string)
	// eksekusi methodS1 yang isinya "wick"
    methodS1.Call(
		// Call(v) v == harus diisi dengan cara []reflect.Value{x}. 
		// x == reflect.ValueOf(z)
		// z == parameter yang mau dimasukin ke func SetStructStudent(name string, age int).

		// []reflect.Value{ reflect.ValueOf(Param_1), reflect.ValueOf(Param_2), ... reflect.ValueOf(Param_n) }

		// Kalau cuma mau isi satu juga bisa.
		// []reflect.Value{ reflect.ValueOf(z), })

		[]reflect.Value{ reflect.ValueOf(z), reflect.ValueOf(2000), 
	})

    fmt.Println("nama :", s1.Name)
    fmt.Println("age :", s1.Age)

}