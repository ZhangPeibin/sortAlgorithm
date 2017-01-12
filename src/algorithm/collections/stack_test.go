package collections

import (
	"testing"
	"fmt"
)

func TestStack(t *testing.T) {
	stack := NewStack()

	stack.Push(&Node{"I"})
	stack.Push(&Node{"Love"})
	stack.Push(&Node{"you"})

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Len())
}
