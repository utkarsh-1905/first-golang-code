package main

import "fmt"

func main() {
	var name string = "Utkarsh"
	fmt.Println("Hello", name)

	age := 19
	fmt.Println("My age is ", age)

	const (
		pi = 3.14
		e  = 2.718
	)
	fmt.Println(pi, e)
}
