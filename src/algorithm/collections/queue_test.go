package collections

import (
	"testing"
	"fmt"
)

func TestNewQueue(t *testing.T) {
	queue := NewQueue()
	queue.Push(&Node{1})
	queue.Push(&Node{2})
	queue.Push(&Node{3})
	queue.Push(&Node{4})

	// 1, 2, 3,4
	fmt.Println(queue.Len())
	queue.rangeQueue()
	fmt.Println()
	fmt.Println(queue.Head())
	fmt.Println(queue.Tail())


	queue.Poll()

	queue.rangeQueue()
	fmt.Println()
	fmt.Println(queue.Head())
	fmt.Println(queue.Tail())

	queue.Poll()

	queue.rangeQueue()
	fmt.Println()
	fmt.Println(queue.Head())
	fmt.Println(queue.Tail())

	queue.Poll()

	queue.rangeQueue()
	fmt.Println()
	fmt.Println(queue.Head())
	fmt.Println(queue.Tail())


	queue.Poll()

	queue.rangeQueue()
	fmt.Println()
	fmt.Println(queue.Head())
	fmt.Println(queue.Tail())

}
