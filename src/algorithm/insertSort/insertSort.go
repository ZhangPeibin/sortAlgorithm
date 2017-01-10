package insertSort

func InsertSort(source []int)  {
	length := len(source)

	if length == 0 || length ==1{
		return
	}

	for i:=0 ; i< length-1 ; i++ {
		for j:=i+1;j > 0 ;j-- {
			if source[j-1] <= source[j] {
				break
			}

			source[j-1],source[j] = source[j],source[j-1]
		}
	}
}