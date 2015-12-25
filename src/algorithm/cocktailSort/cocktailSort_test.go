//cocktailSort_test.go
package cocktailSort

import (
	"testing"
)

func TestCocktailSort(t *testing.T) {
	values := []int{2, 1, 5, 3, 6}

	CocktailSort(values)

	if values[0] != 1 || values[4] != 6 {
		t.Error("CocktailSort() failed", values, "Excepted values 1,2,3,5,6")
	}

}
