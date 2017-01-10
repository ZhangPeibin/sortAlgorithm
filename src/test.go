package main

import "fmt"

func main() {
	a := make([]int,5,10)
        a = append(a,1)
	a = append(a,1)
	a = append(a,1)
	a = append(a,1)
	a = append(a,1)
	a = append(a,1)

	fmt.Println(len(a))
	fmt.Println(cap(a))

}
