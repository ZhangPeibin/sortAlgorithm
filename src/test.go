package main

import "fmt"

var a []int

func main() {
	//a := make([]int,0)
        a = append(a,1)
	a = append(a,2)
	a = append(a,3)
	a = append(a,4)
	a = append(a,5)

	//for _,value := range a{
		//fmt.Println(value)
	//}

	b := a[0:5]

	for _,value := range b{
		fmt.Println(value)
	}
}
