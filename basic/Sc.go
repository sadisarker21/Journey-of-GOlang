package main

import (
	"fmt"
)

var (
	a = 12
	b = 14
	c = 10
)

func printf(num int) {
	fmt.Println(num)
}

func add(x int, y int) {
	res := x + y
	printf(res)
}

func biger_of_three_number(num1 int, num2 int, num3 int) {
	if num1 > num2 && num1 > num3 {
		printf(num1)
	} else if num2 > num1 && num2 > num3 {
		printf(num2)
	} else {
		printf(num3)
	}

}

func main() {
	add(a, b)
	biger_of_three_number(a, b, c)
}
