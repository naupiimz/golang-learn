package main

import "fmt"

func main() {

	// title := "golang is the best"

	// for _, val := range title {
	// 	switch string(val) {
	// 	case "a", "i", "u", "e", "o":
	// 		fmt.Println(string(val))
	// 	}
	// }

	// var languages [5]string

	// languages[0] = "GO"
	// languages[1] = "Gin"
	// languages[2] = "BeeGO"
	// languages[3] = "Echo"
	// languages[4] = "Tbd."

	// fmt.Println(languages)
	// fmt.Println(len(languages))

	// var consolesG []string

	// consolesG = append(consolesG, "Xbox360")
	// consolesG = append(consolesG, "Switch Pro")

	// fmt.Println(consolesG)

	// myMap := map[string]string{}

	// myMap["adam"] = "adam"

	// fmt.Println(myMap)

	// value, isExist := myMap["adam"]

	// if isExist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("doesnt exist")
	// }

	// students := []map[string]string{
	// 	{"name": "upi", "score": "good"},
	// 	{"name": "ray", "score": "bad"},
	// }

	// for _, v := range students {
	// 	fmt.Println(v["name"])
	// 	fmt.Println(v["score"])
	// }

	// scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	// goodScores := []int{}

	// i := 0

	// for _, v := range scores {
	// 	i = i + v

	// 	if v >= 90 {
	// 		goodScores = append(goodScores, v)
	// 	}
	// }

	// result := float64(i) / float64(len(scores))

	// fmt.Println(result)
	// fmt.Println(goodScores)

	numarr := []int{2, 3, 4, 5, 6}

	result := calculate(numarr)
	fmt.Println(result)
}

func calculate(nums []int) int {
	total := 0
	for _, num := range nums {
		total = total + num
	}

	return total
}
