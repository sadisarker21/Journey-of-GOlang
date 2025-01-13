package main

import (
	"fmt"
)

func messager() {
	fmt.Println("Hellow go !")
}

func gender(fg string){
	if fg == "male"{
		fmt.Println("you are man")
	} else{
		fmt.Println("you are male")
	}

}
func main() {
	gn:="male"
	gender(gn)
	messager()
}
