package main

import "fmt"

func removeBack(array []int) ([]int, int) {

	return array[:len(array)-1], array[len(array)-1]
}

func removeFront(array []int) ([]int, int) {
	return array[1:], array[0]
}

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	var removedNumber int

	for i := 0; i < 5; i++ {

		a, removedNumber = removeBack(a)
		fmt.Println(removedNumber)

	}

	fmt.Println(a)

	for i := 0; i < 5; i++ {
		b, removedNumber = removeFront(b)
		fmt.Println(removedNumber)
	}

	fmt.Println(b)
	//정확히 말해본다면 제거가 아니고 원하는 배열의 주소값을 이용하여 다시 재지정해준것이다.
}
