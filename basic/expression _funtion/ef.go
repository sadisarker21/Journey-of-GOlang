package main

import (
	"fmt"
)

func main() {

//expression funtion
	add := func(a, b int) {
		c := a + b
		fmt.Println(c)
	}

	add(3, 4)
}
