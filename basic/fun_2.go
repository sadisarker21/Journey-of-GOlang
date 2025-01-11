package main

import "fmt"
func iseven(num int){
if num%2==0{
	fmt.Println("This is a even number")
} else {
	fmt.Println("this is odd number")
}
}

func main() {
a := 14

iseven(a)
}
