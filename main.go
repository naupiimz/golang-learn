package main

import "fmt"

func main() {

	title := "golang is the best"

	for _, val := range title {
		switch string(val) {
		case "a", "i", "u", "e", "o":
			fmt.Println(string(val))
		}
	}

}
