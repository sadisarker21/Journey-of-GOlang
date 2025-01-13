package main

import "fmt"

func messageprint(){
	fmt.Println("Welcome to ours application !")
}
func userinput()string{
	fmt.Println("Enter your name :")
	var name string
	fmt.Scan(&name)
	return name
}

func Tow_number_input()(int,int){
	var num_1 int
	var num_2 int
	fmt.Println("Please enter your first number :")
	fmt.Scanln(&num_1)
	fmt.Println("Please enter your secand number ")
	fmt.Scanln(&num_2)
	return num_1,num_2

}
func add(num_1 int,num_2 int)int{
	sum := num_1 + num_2
	return sum
}

func disply( n string,s int){
	fmt.Println("Hellow !", n)
	fmt.Println("Samation is :", s)
	
}

func goodbye(){
	fmt.Println("Thanks for using ours application")
	fmt.Println("Goodbye")
}


func main() {
messageprint()
name :=userinput()
num1,num2 :=Tow_number_input()	
sum :=add(num1,num2)
disply(name,sum)
goodbye()


}
