package main

import "fmt"
var a = 33
var b = 44

func sum(x int, y int)  {
	z := x + y
	fmt.Println(z)


}

func main() {

	p := 23
	q := 11

	sum(p, q)
	sum(a,b)
	sum(p,a)
	sum(q,b)
	sum(a,z) //scope nai 

}
