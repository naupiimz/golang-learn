package main

import (
	"fmt"
	"golang-learn/calculation"
	"golang-learn/multiplier"
)

func main() {
	fmt.Println("reset")

	result := calculation.Add(12312412, 19484124)

	result2 := multiplier.Multiply(12123, 124499)

	fmt.Println(result)

	fmt.Println(result2)
}
