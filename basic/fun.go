package main

import "fmt"

func add(num_1 int, num_2 int) {
	sum := num_1 + num_2
	fmt.Println(sum)
}

func main() {
	a := 12
	b := 12
	add(a, b)
}
