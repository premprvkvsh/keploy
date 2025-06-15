package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	myNum := 23
	var ptr = &myNum

	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("Value of pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNum)
}
