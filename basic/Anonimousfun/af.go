package main

import "fmt"

// Anonimous funtion
// Imediately Invoke Funtions expression
func main() {
	func(a, b int) {
		c := a + b
		fmt.Println(c)
	}(3, 5)

}

func init() {
	fmt.Println("I am calling First")
}
