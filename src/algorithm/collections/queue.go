package collections

import (
	"sync"
)


//栈
type Queue struct {
	sync.RWMutex
	elementData []*Node
	elementCount int
	capacityIncrement int
}