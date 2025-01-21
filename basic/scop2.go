package main

import "fmt"

func main() {
    x := 10
    {
        y := 20
        fmt.Println(x, y) // Accessible here
    }
    // fmt.Println(y) // Error: y is undefined here
}
