package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	fmt.Println("Recursion")
	fmt.Println("=========")

	fmt.Println(factorial(7))

	var fib func(n int) int

	var myfib int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	myfib = fib(7)

	fmt.Println(myfib)

	myfib = fib(10)

	fmt.Println(myfib)
}
