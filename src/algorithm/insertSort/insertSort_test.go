package insertSort

import (
	"testing"
	"fmt"
)

func TestInsertSort(t *testing.T) {
	values := []int{2, 1, 5, 3, 6}
	InsertSort(values)
	for i:=0;i<len(values) ;i++  {
		fmt.Println(values[i])
	}
}
