package main

import (
	"fmt"
)

func main() {
	gender := "e"
	switch gender {
	case "M", "m":
		fmt.Println("He is a boy")

	case "F", "f":
		fmt.Println("She is a girl")
	default:
		fmt.Println("You are gay")
	}
}
