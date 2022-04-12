package main

import (
	"fmt"
	// "github.com/mhaikalkal/learnGO/condition"
	// "github.com/mhaikalkal/learnGO/iteration"
	// "github.com/mhaikalkal/learnGO/variable"
	// "github.com/mhaikalkal/learnGO/array"
	// "github.com/mhaikalkal/learnGO/slice"
	// mapping "github.com/mhaikalkal/learnGO/map" // memakai Alias.
	// jadi pas manggil salah satu func pointer di file main.go, gausah pointer.NamaFunc. tapi p.NamaFunc
	// p "github.com/mhaikalkal/learnGO/pointer"

	/*
		Kadang kita hanya ingin menjalankan init function di package tanpa harus mengeksekusi salah satu function yang ada di package Secara default,
		Go akan komplain ketika ada package yang di import namun tidak digunakan Untuk menangani hal tersebut,
		kita bisa menggunakan blank identifier (_) sebelum nama package ketika melakukan import.
	*/

	"github.com/mhaikalkal/learnGO/imexport"
)

func text() {
	fmt.Println("Hello, Haikal")
}

func main() {
	// text()
	// variable.VarFunc()
	// variable.TugasJCC_2()
	// condition.IfCondition()
	// condition.SwitchCase()
	// iteration.ForIteration()
	// array.ArrayFunc()
	// array.ArrayFuncWithIteration()
	// array.ArrayMultidomension()
	// slice.SliceFunc()
	// slice.SliceWithIteration()
	// slice.RNGPercentBased()
	// mapping.MapFunc()
	// p.PointerTest()

	/*
		karena ImEx_port.go satu package / folder dengan main.go
		jika kita cuma `go run main.go` func dari ImEx gabisa dipakai di main.go
		maka untuk eksekusi keduanya bisa pake `go run main.go ImEx_port.go` atau `go run .`
		dan juga untuk memanggil function dari ImEx yang hurufnya kecil pun possible. karena scope keduanya masih local (1 neighborhood)
	*/

	SayHello("Haikal")  
	fmt.Println(calculate(9, 11)) 

	fmt.Println(imexport.DatabaseConnection)
	
}