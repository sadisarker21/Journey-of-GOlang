package main
import "fmt"
func main(){
	var num int
	fmt.Println("Enter a number:")
	fmt.Scan(&num)
	if num%2 ==0{
		fmt.Println("The number is even")

	}else {
		fmt.Println("the number is odd")
	}
}