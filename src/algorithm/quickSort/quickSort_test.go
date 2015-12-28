package quickSort

import "testing"

func TestQuick(t *testing.T) {
	slice := []int{0, 3, 1, 412, 13}

	QuickSort(slice)

	if slice[0] != 0 || slice[1] != 1 || slice[2] != 3 || slice[3] != 13 {
		t.Error("Quick sort failed")
	}

}
