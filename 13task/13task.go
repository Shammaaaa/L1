package main

import "fmt"

func main() {
	a, b := 10, 20
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}
