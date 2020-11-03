package main

import "fmt"

func main() {
	x := 1
	y := 2

	swap(&x, &y)

	fmt.Println(x)
	fmt.Println(y)

}

func swap(xPtr *int, yPtr *int) {
	var temp int
	temp = *xPtr
	*xPtr = *yPtr
	*yPtr = temp

}
