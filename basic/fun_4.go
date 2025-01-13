package main

import "fmt"

func fact(n int) int {

	
	F := n * (n - 1)
	return F
}

func main() {
	NUM := 12
	Result := fact(NUM)
	fmt.Println("The factorial is :", Result)
}
