package condition

import "fmt"

func IfCondition() {
	
	var nilai int = 7

	if nilai >= 10 {
		fmt.Println("You got", nilai ,"You have a perfect score")
	} else if nilai <= 9 && nilai >= 7  {
		fmt.Println("You got", nilai ,"You have a great score")
	} else if nilai <= 6 {
		fmt.Println("You got", nilai ,"You have to learn more")
	}

	var poin float32 = 36

	if score := (poin/5)*10; score >= 90 {
		fmt.Println("Your score is", score)
		fmt.Println("You got an A")
	} else if score < 90 && score >= 70 {
		fmt.Println("Your score is", score)
		fmt.Println("You got a B")
	} else if score < 70 && score >= 60 {
		fmt.Println("Your score is", score)
		fmt.Println("You got a C")
	} else {
		fmt.Println("Your score is", score)
		fmt.Println("You got a D")
	}

}

func SwitchCase() {
	// switch di GO, akan berhenti ketika udah masuk di case yang sesuai.

	var nilaiX = 6
	
	switch nilaiX {
	case 10:
		fmt.Println("Perfect Score")
	case 8:
		fmt.Println("Great Score")
	case 7, 6, 5, 4:
		fmt.Println("Good Score")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("You can do better")
			// kalau mau multi output pakai kurawal
		}
	}

	// switch case gaya if-else
	var pointY = 6

	switch {
	case pointY == 8:
		fmt.Println("perfect")
	case (pointY < 8) && (pointY > 3):
		fmt.Printf("awesome point, you got %d PointY %s", pointY, "\n")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}


	// fallthrough, meski udah masuk case yang sesuai dan di run, dia bakal masuk ke case selanjutnya.
	var pointX = 6

	// pada kasus ini si pointX masuk ke case yang kedua, output = awesome, karena ada fallthrough dia juga masuk case ketiga.
	// jadi output finalnya = awesome, you need to learn more
	switch {
	case pointX == 8:
		fmt.Println("perfect")
	case (pointX < 8) && (pointX > 3):
		fmt.Println("awesome")
		fallthrough
	case pointX < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	var pointZ = 10
	
	if pointZ > 7 {
		switch pointZ {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if pointZ == 5 {
			fmt.Println("not bad")
		} else if pointZ == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if pointZ == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}