package main

import (
	"fmt"

	"github.com/bryan01-Martin/dataStruct"
)

func main() {
	list := &dataStruct.LinkedList{}
	list.AddNode(0)
	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}
	list.PrintNodes()
	list.PrintReverse()

	list.RemoveNode(list.Root.Next)
	list.PrintNodes()
	list.PrintReverse()

	list.RemoveNode(list.Root)
	list.PrintNodes()
	list.PrintReverse()

	list.RemoveNode(list.Tail)
	list.PrintNodes()
	list.PrintReverse()
	fmt.Printf("tail : %d\n", list.Tail.Val)

	stack := []int{}

	for i := 1; i < 10; i++ {
		stack = append(stack, i)
	}
	fmt.Println(stack)

	for len(stack) > 0 {
		var last int
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]
		fmt.Println(last)
	}

	queue := []int{}

	for i := 1; i < 5; i++ {
		queue = append(queue, i)
	}
	fmt.Println(queue)
	for len(queue) > 0 {
		var front int
		front, queue = queue[0], queue[1:]
		fmt.Println(front)
	}

	stack2 := dataStruct.NewStack()

	for i := 1; i <= 5; i++ {
		stack2.Push(i)
	}

	fmt.Println("New Stack!!")
	for !stack2.Empty() {
		val := stack2.Pop()
		fmt.Printf("%d -> ", val)
	}
	fmt.Println("")

	queue2 := dataStruct.NewQueue()

	for i := 1; i <= 5; i++ {
		queue2.Push(i)
	}

	fmt.Println("New Queue!!")
	for !queue2.Empty() {
		val := queue2.Pop()
		fmt.Printf("%d -> ", val)
	}
	fmt.Println("")

}
