package main

import "fmt"

func min_convert_sec(){
	var m int
	fmt.Println("Enter your minute :")
	fmt.Scan(&m)
	if m <=60{
		fmt.Println("The secand is : ",m*60)
	}else {
		fmt.Println("Invalid number")
	}
	
}

func main() {
min_convert_sec()	
}
