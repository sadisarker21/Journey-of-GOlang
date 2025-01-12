package main

import "fmt"

func sendMeassage(name string) string {
	return name
}
func numbers(num_1 int, num_2 int) (int, int, int) {
	mul := num_1 * num_2
	min := num_1 - num_2
	div := num_1 + num_2

	return mul, min, div
}
func main() {
	name := "HI GOlnag"
	a := 12
	b := 14

	surname := sendMeassage(name)
	m,M,d := numbers(a, b)
	fmt.Println(m)
	fmt.Println(M)
	fmt.Println(d)
	fmt.Println(surname)
}
