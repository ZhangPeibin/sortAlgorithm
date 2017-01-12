package collections

import (
	"errors"
	"sync"
)



//栈
type Stack struct {
	sync.RWMutex
	elementData []*Node
	elementCount int
	capacityIncrement int
}

func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(object *Node) error {

	stack.Lock()
	defer stack.Unlock()

	if object == nil {
		 return errors.New("can not add a nil interface{}")
	}

	stack.elementData = append(stack.elementData,object)

	stack.elementCount++

	return nil
}

func (stack *Stack) Pop()(object *Node) {

	stack.Lock()
	defer stack.Unlock()

	if stack.elementCount == 0 {
		return nil
	}

	object = stack.elementData[stack.elementCount-1]

	stack.elementData = stack.elementData[:stack.elementCount-1]
	stack.elementCount--
	return
}

func (stack *Stack) Peek()(object *Node) {
	stack.Lock()
	defer stack.Unlock()

	if stack.elementCount == 0{
		return nil
	}

	object = stack.elementData[stack.elementCount-1]

	return
}

func (stack *Stack) Len()int{
	return stack.elementCount
}