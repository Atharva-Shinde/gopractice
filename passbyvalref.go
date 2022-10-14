package main

import "fmt"

func main() {
	//pass by value
	a := 1
	fmt.Print("pass by value\n")
	fmt.Printf("value of a: %v, memory location of a: %p\n", a, &a)
	add(a)
	fmt.Printf("main function after execution of add func value of a: %v\n", a)
	fmt.Print("=================================================\n")

	//pass by reference
	b := 10
	fmt.Print("pass by reference\n")
	fmt.Printf("value of b: %v, memory location of b: %p\n", b, &b)
	addptr(&b)
	fmt.Printf("main function after execution of addptr func value of b: %v\n", b)
}

// pass by value
func add(x int) {
	fmt.Printf("inside add func before operation, value of x: %v, memory location of x: %p\n", x, &x)
	x++
	fmt.Printf("inside add func after operation, value of x: %v, memory location of x: %p\n", x, &x)
}

// pass by reference
func addptr(y *int) {
	fmt.Printf("inside addptr func before operation, value of *y: %v, value of y %v, memory location of y: %p\n", *y, y, &y)
	*y++
	fmt.Printf("inside addptr func after operation, value of *y: %v, value of y %v, memory location of y: %p\n", *y, y, &y)
}
