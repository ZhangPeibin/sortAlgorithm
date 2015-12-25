package bubbleSort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	values := []int{2, 1, 5, 3, 6}

	BubbleSort(values)

	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 || values[4] != 6 {
		t.Error("BubbleSort() failed.Got", values, "Expected 1,2,3,5,6")
	}

}
