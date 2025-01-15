package main

import (
	"fmt"
)

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func main() {
	input := "Hello, World!"
	reversed := reverseString(input)
	fmt.Println("Original String:", input)
	fmt.Println("Reversed String:", reversed)
}
