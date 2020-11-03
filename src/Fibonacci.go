package main

import "fmt"

func main() {

	fmt.Println("Enter int")
	var num int

	fmt.Scanln(&num)
	fmt.Println(fib(num))

}

func fib(num int) uint {

	var total uint = 0
	if num == 0 {

		return total

	} else if num == 1 {

		return total + 1
	}

	return fib(num-1) + fib(num-2)
}
