package main
import "fmt"

func swap(n1 int,n2 int)(int,int){
	var tem int
	tem =n1
	n1 =n2
	n2 =tem
	fmt.Println(n1,n2)
	return n1,n2

}

func main(){
	a:=2
	b:=3
 swap(a,b)	
}