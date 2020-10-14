package main

import (
	"fmt"
)

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func main() {
	fmt.Println(Add(1, 2))
	fmt.Println(Subtract(1, 2))
	fmt.Println(Multiply(1, 2))
}
