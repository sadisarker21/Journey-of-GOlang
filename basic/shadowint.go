package main

import "fmt"

func main() {
	a := 23

	age := 12

	if age > 10 {
		a := 10
		fmt.Println(a)
	}

	fmt.Println(a)
}
