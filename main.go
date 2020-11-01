package main

import "fmt"

func main() {
	fmt.Println(findTheMostInArgs(1, 2, 3, 4, 5))

}

func findTheMostInArgs(args ...int) int {

	theMost := 0

	for _, v := range args {
		if theMost < v {
			theMost = v
		}

	}
	return theMost
}
