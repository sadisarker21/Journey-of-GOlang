package main

import "fmt"

func main() {
	number := 13
	if number > 0 {
		fmt.Println("This is positive number", number)
	} else if number < 0 {
		fmt.Println("this is negitive number",number)
	} else {
		fmt.Println("The number is zero.")
	}
}
