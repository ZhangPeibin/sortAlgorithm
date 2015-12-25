package cocktailSort

//鸡尾酒排序
//也叫做定向冒泡排序算法，鸡尾酒搅拌排序，搅拌排序
//其效率相比较冒泡排序稍微的好一点，但跟冒泡算法一样差劲
func CocktailSort(values []int) {

	leftIndex := 0
	rightIndex := len(values) - 1

	for leftIndex < rightIndex {
		for index := leftIndex; index < rightIndex; index++ {
			if values[index] > values[index+1] {
				values[index], values[index+1] = values[index+1], values[index]
			}
		}

		rightIndex--

		for index := rightIndex; index > leftIndex; index-- {
			if values[index] < values[index-1] {
				values[index], values[index-1] = values[index-1], values[index]
			}
		}

		leftIndex++

	}
}
