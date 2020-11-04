package main

import (
	"datastruct"
	"fmt"
)

func main() {
	stack := []int{}

	for i := 1; i < 5; i++ {
		stack = append(stack, i)

	}
	fmt.Println(stack)
	var last int

	for len(stack) > 0 {
		last = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		fmt.Println(last)
	}

	fmt.Println()

	queue := []int{}

	for i := 1; i < 5; i++ {
		queue = append(queue, i)

	}

	fmt.Println(queue)

	var first int

	for len(queue) > 0 {
		first = queue[0]
		queue = queue[1:]

		fmt.Println(first)

	}

	fmt.Println("New Stack")

	stack2 := datastruct.NewStack()

	for i := 1; i <= 5; i++ {
		stack2.Push(i)
	}

	for !stack2.Empty() {
		val := stack2.Pop()
		fmt.Printf("%d ->", val)

	}
	fmt.Println()


	fmt.Println("New Queue")


	queue2 := datastruct.NewQueue()

	for i := 1; i <=5; i++ {
		queue2.Push(i)
	}

	for !queue2.Empty() {
		val := queue2.Pop()
		fmt.Printf("%d ->", val)
	}


}
