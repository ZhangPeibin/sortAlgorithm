package bubbleSort

import (
	"fmt"
)

//冒泡排序的实现
//冒泡的原理是比较相邻两数的大小,如果前一个数字比后一个大,则交换
//每一次循环会把最大的一个数字放到最后一位
//所以得循环n-1次
func BubbleSort(source []int) {

	//用来做循环完成标志位,当某次最外层循环
	//并没有导致交换,那么就直接跳出循环
	//因为没有交换,就代表数组已经不需要排序了
	finishSort := true

	fmt.Println("开始排序...")
	//最外层的循环,最多循环len-1次
	for i := 0; i < len(source); i++ {
		finishSort = true //循环开始前设置标志位为true
		fmt.Println("---开始一次外部循环...")
		//为什么要循环len-1-i次?
		//因为每一次的循环都会把最大的数放到最后
		//所以,每循环n次，那么最后n个数就不会排序
		for j := 0; j < len(source)-1-i; j++ {
			fmt.Println("------开始一次内部循环...")
			if source[j] > source[j+1] {
				source[j], source[j+1] = source[j+1], source[j]
				finishSort = false
			}
		}

		if finishSort == true {
			fmt.Println("排序完成...")
			break
		}
	}
}
