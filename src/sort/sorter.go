package main

import (
	"../algorithm/bubbleSort"
	"../algorithm/cocktailSort"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File contains value after sorted")
var algorithm *string = flag.String("a", "algorithm", "algorithm to sort")

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("inflie = ", *infile, ",outfile=", *outfile, ",algorithm=", *algorithm)
	}

	values, err := readFile(*infile)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("start sort :", values)
	}

	fmt.Println("sort algorithm", *algorithm)

	switch *algorithm {
	case "bubble":
		bubbleSort.BubbleSort(values)
	case "cocktail":
		cocktailSort.CocktailSort(values)
	}

	writeValues(values, *outfile)

	fmt.Println("end sort :", values)

}

func readFile(path string) (values []int, err error) {
	file, err := os.Open(path)

	fmt.Println("file name :", file.Name())

	if err != nil { //如果error不为null
		fmt.Println("open file with", path, "failed")
		return
	}

	//延迟执行该语句,用来关闭file文件流
	defer file.Close()

	reader := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line, isPrefix, readErr := reader.ReadLine()

		if readErr != nil {
			if readErr != io.EOF {
				err = readErr
			}
			break
		}

		if isPrefix {
			fmt.Println("a too long line,seems unexcepted.")
			return
		}

		//将我们的字节转为字符串
		str := string(line)

		//将字符串转为我们的int类型书籍
		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}

	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("fail to create a out file ", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return nil
}
