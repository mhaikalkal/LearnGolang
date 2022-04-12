package iteration

import "fmt"

func ForIteration() {
	// Perulangan di bahasa go cuma ada for saja.
	for i := 0; i < 5; i++ {
		fmt.Println("I Number", i)

	}

	// Cara penulisan for ada banyak di bahasa Go
	var j = 5
	for j < 10{
		fmt.Println("J Number", j)
		j++

	}

	// Bisa di berhentiin pake break juga
	var k = 10
	for {
		fmt.Println("K Number", k)
		k++

		if k == 15 {
			break
		}
	}

	// break continue juga bisa
	for l := 0; l <= 10; l++{
		if l%2 == 1 { // kalau l/2 sisa 1, maka print `Hasil ganjil` kalau l/2 = 0, berarti genap
			fmt.Println("L Number", l)
			continue
		}

		if l >= 10 {
			break
		}
	}


}

func GambarIteration() {
	n := 7
	for i := n; i > 0; i--{
		for j := i; j < n; j++{
			fmt.Print("#")
		}
		fmt.Println("#")
	}

}