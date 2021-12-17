package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println("")
	variabel()
	fmt.Println("")
	angka()
	fmt.Println("")
	logika()
	fmt.Println("")
	kata()
	fmt.Println("")
	konstanta()
	fmt.Println("")
	jenisoperator()
	fmt.Println("")
	kondisi()
}

func variabel() {
	//manifest typing variable declaration
	//deklarasi variabel menggunakan keyword var
	var firstName string = "nau"

	var lastName string
	lastName = "piimz"

	fmt.Printf("Hello %s %s! \n", firstName, lastName)

	//type inference variable declaration
	var firstTitle string = "Software Dev"
	lastTitle := "Web Dev"
	fmt.Printf("%s, %s\n", firstTitle, lastTitle)

	//var keyword tanpa tipe data
	var room = "XI-III"
	fmt.Println(room)
}

func angka() {
	var postNum uint8 = 89
	var negNum = -1234567890

	fmt.Printf("bilangan positif: %d\n", postNum)
	fmt.Printf("bilangan negati: %d\n", negNum)

	//angka numerik
	var decimalNum = 2.62
	fmt.Printf("bilangan desimal: %f\n", decimalNum)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNum)
}

func logika() {
	var exist bool = true
	fmt.Printf("exist? %t\n", exist)
}

func kata() {
	var message string = "Halo"
	//contoh penggunaan backtick
	var message2 string = `Halo
salam`

	fmt.Printf("message: %s\n", message)
	fmt.Printf("message2: %s\n", message2)
}

func konstanta() {
	const firstName string = "nau"
	fmt.Print("halo ", firstName, "!\n")
}

func jenisoperator() {
	//aritmatika
	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	//logika
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}

func kondisi() {
	var point = 1

	if point == 10 {
		fmt.Println("perfect")
	} else if point > 5 {
		fmt.Println("passed")
	} else if point == 4 {
		fmt.Println("so close")
	} else {
		fmt.Println("not passed")
	}

	var score = 8840.0

	if percent := score / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good!\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad!\n", percent, "%")
	}

	//switch case, dapat menampung lebih dari 1 kondisi
	var result = 6

	switch result {
	case 8:
		fmt.Printf("perfect")
	case 7:
		fmt.Printf("awesome")
	default:
		fmt.Printf("not bad")
	}
}
