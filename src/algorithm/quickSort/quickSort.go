package quickSort

import (
	"fmt"
)

//快速排序算法的实现
//快速排序的原理
//指定一个数作为标示
//依次比较整个数组,将比标示的数小的放到其左边
//比标示大的数放在其右边
//最核心的，就是使用了‘分冶法’

//根据参数可以知道传递的是slice而不是数组
//所以操作slice是直接操作内存
func QuickSort(source []int) {

	fmt.Println("start sort ... ")

	//数组长度小于等于1的时候是不需要排序的
	if len(source) <= 1 {
		return
	}

	//设置基准为第一个数字
	mid, i := source[0], 1

	//设置左右的下标
	head, tail := 0, len(source)-1

	//做循环
	for head < tail {

		//如果第i个数大于我们的基准
		if source[i] > mid {
			//那么将这个数与最后一个数字替换
			//替换的原因
			//每次替换后最后一个就不需要在比较
			//而原来这个位置的数据就会在下一次循环中
			//跟mid做比较,也就是用倒叙的方式从数组的后面
			//跟mid做比较
			source[i], source[tail] = source[tail], source[i]
			tail--
		} else {
			//假设数据小于mid
			//直接讲该数据放到第head上面去
			//该循环跟上面的循环是一个道理
			//只不过一个从数组后面循环
			//而这个得从数组前面开始循环
			source[head], source[i] = source[i], source[head]
			head++
			i++
		}
	}

	source[head] = mid
	QuickSort(source[:head])
	QuickSort(source[head+1:])
}
