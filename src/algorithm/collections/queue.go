package collections

import (
	"sync"
	"fmt"
)


//栈
type Queue struct {
	lock *sync.RWMutex
	elementData []*Node
	elementCount int
	head *Node
	tail *Node
}

func NewQueue() (*Queue) {
	return &Queue{lock:&sync.RWMutex{}}
}

func (queue *Queue)Push(node *Node)  {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	queue.elementData = append(queue.elementData,node)
	queue.elementCount++

	if queue.head == nil{
		queue.head = node
	}

	queue.tail = queue.elementData[queue.elementCount-1]
}

//get a node from queue,and remove it in queue
func (queue *Queue)Poll()(*Node){
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.elementCount == 0 {
		return nil
	}

	node := queue.elementData[0]

	queue.elementData=queue.elementData[1:queue.elementCount]

	queue.elementCount--

	if queue.elementCount == 0 {
		queue.head = nil
		queue.tail = nil
	}else{
		queue.head = queue.elementData[0]
		queue.tail = queue.elementData[queue.elementCount-1]
	}

	return node

}

func (queue *Queue)Len() int {
	return queue.elementCount
}

func (queue *Queue)Head() (*Node) {
	return queue.head
}

func (queue *Queue)Tail() (*Node) {
	return queue.tail
}

func (queue *Queue)rangeQueue() {
	for _,value :=range queue.elementData{
		fmt.Print(value )
		fmt.Print(",")
	}
}