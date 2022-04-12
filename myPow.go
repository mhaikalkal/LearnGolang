package main

import "fmt"

func main() {
	fmt.Println(myPow(2, 10))
	fmt.Println(myPow(2, 0))
	fmt.Println(myPow(2, -2))

}

func myPow(x float64, n int) float64 {
	var sliceX []float64
	var sum float64
	var z float64 = 1

	if n > 0 {
		for i := 0; i < n; i++ {
			sliceX = append(sliceX, x)
		}
		for _, val := range sliceX {
			z *= val
			sum = z
		}
	} else if n < 0 {
		for i:= 0; i < -(n); i++ {
			sliceX = append(sliceX, x)
		}

		for _, val := range sliceX {			
			z *= val
		}
		
		sum = 1 / z
	} else if n == 0 {
		sum = 1
	}

	return sum

}