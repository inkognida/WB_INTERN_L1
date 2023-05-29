package main

import "fmt"

func main() {
	a := 10
	b := 11

	fmt.Println(a,b)

	a,b = b,a
	fmt.Println(a,b)
}