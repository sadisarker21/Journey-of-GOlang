package main

import "fmt"

func watermalon_devide(w int) bool {
	if w > 2 && w%2 == 0 {
		return true
	}
	return false
}

func main() {
	var watermalon int
	fmt.Println("Enter the watermalon weight : ")
	fmt.Scan(&watermalon)



	w:=watermalon_devide(watermalon)
	fmt.Println(w)
}
