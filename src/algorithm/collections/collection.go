package collections

import "fmt"

type Collections interface {
	Add(a interface{})
}


type Node struct{
	node interface{}
}

func (node *Node)String()string  {
	return fmt.Sprint(node.node)
}
