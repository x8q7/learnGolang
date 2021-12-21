package main

import "fmt"

func InsertSort(arr []int) []int {

	length := len(arr)
	if length == 0 {
		return []int{}
	}
	if length == 1 {
		return arr
	}
	for i := 1; i < length; i++ {
		cur := arr[i]
		preIndex := i - 1

		for preIndex >= 0 && cur < arr[preIndex] {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}

		if preIndex+1 != i {
			arr[preIndex+1] = cur
		}

	}
	return arr
}

func main() {
	arr := []int{2, 42, 5, 627, 4, 88, 92, 2, 5, 8}

	sortArr := InsertSort(arr)

	fmt.Println(sortArr)
}
