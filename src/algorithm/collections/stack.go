package collections

import "errors"



type Stack struct {
	elementData []interface{}
	elementCount int
	capacityIncrement int
}

func Stack() *Stack {
	return StackCapac(5)
}

func StackCapac(capacityIncrement int) *Stack {
	return &Stack{
		elementCount:0,
		elementData:make([]interface{},capacityIncrement),
		capacityIncrement:capacityIncrement,
	}
}

func (stack *Stack) Push(object interface{}) error {
	if object == nil {
		 return errors.New("can not add a nil interface{}")
	}

	stack.resizeIfNeed()

	stack.elementData = append(stack.elementData,object)

	stack.elementCount++

	return nil
}

func (stack *Stack) Pop()(object interface{}) {

	object = stack.elementData[stack.elementCount-1]

	stack.elementData = stack.elementData[:stack.elementCount-1]

	return
}


func (stack *Stack) resizeIfNeed() {
	if ( cap(stack.elementData) == stack.elementCount+2) {
		t := make([]interface{},len(stack.elementData),cap(stack.elementData)+stack.capacityIncrement*2+1)
		copy(t,stack.elementData)
		stack.elementData = t
	}
}