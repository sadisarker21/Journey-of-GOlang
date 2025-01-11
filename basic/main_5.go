package main

import "fmt"

func main() {
	viowel := "y"

	switch viowel {
	case "a", "e", "i", "o", "m":
		fmt.Println("This is a alphabate")
	default:
		fmt.Println("This is a consonent")
	}

}
