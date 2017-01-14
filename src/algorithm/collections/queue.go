package collections

import (
	"sync"
	"fmt"
)

type Queue interface {
	//return the length of this queue
	Len() int

	//push a node to queue
	Add(*Node) bool

	//remove and get a node from queue,if queue empty ,it will throw a error
	Remove() (error,*Node)

	//Retrieves a node from queue ,if queue empty,it will return nil
	Poll() *Node

	//Retrieve the head from queue,it queue empty ,it will return error
	Element() (error,*Node)

	//Retrieve the head from queue,it queue empty ,it will return nil
	Peek() *Node
}

//栈
type LinkedList struct {
	lock         *sync.RWMutex
	elementData  []*Node
	elementCount int
	head         *Node
	tail         *Node
}

func NewQueue() (*Queue) {
	return &Queue{lock:&sync.RWMutex{}}
}

func (queue *Queue)Push(node *Node) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	queue.elementData = append(queue.elementData, node)
	queue.elementCount++

	if queue.head == nil {
		queue.head = node
	}

	queue.tail = queue.elementData[queue.elementCount - 1]
}

//get a node from queue,and remove it in queue
func (queue *Queue)Poll() (*Node) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.elementCount == 0 {
		return nil
	}

	node := queue.elementData[0]

	queue.elementData = queue.elementData[1:queue.elementCount]

	queue.elementCount--

	if queue.elementCount == 0 {
		queue.head = nil
		queue.tail = nil
	} else {
		queue.head = queue.elementData[0]
		queue.tail = queue.elementData[queue.elementCount - 1]
	}

	return node

}

/**
* Retrieves, but does not remove, the head of this queue.  This method
* differs from {@link #peek peek} only in that it throws an exception
* if this queue is empty.
*
* @return the head of this queue
* @throws NoSuchElementError if this queue is empty
*/
func (queue *Queue)Element() (err error, node *Node) {

}

/**
* Retrieves, but does not remove, the head of this queue,
* or returns {@code null} if this queue is empty.
*
* @return the head of this queue, or {@code null} if this queue is empty
*/
func (queue *Queue)Peek() (node *Node) {

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
	for _, value := range queue.elementData {
		fmt.Print(value)
		fmt.Print(",")
	}
}