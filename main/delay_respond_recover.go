package main

import "fmt"

func main() {

	defer func() {

		var scanVar int
		fmt.Scanf("%d", scanVar)
		str := recover()
		fmt.Println(str)

	}()
	panic("고장!")
}
