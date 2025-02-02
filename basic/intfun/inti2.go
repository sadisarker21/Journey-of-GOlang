package main
import "fmt"
var a=23
func main(){
	fmt.Println(a)
}

func init(){
	fmt.Println(a)
	a=10
}

