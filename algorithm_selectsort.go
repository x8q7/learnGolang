package main

import "fmt"

func SelectSort(arr []int) []int {

	length := len(arr)
	if length == 0 {
		return []int{}
	}
	if length == 1 {
		return arr
	}
	for i := 0; i < length; i++ {
		min := arr[i]
		minIndex := i
		for j := i + 1; j < length; j++ {
			if min > arr[j] {
				min = arr[j]
				minIndex = j
			}
		}
		if i != min {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	return arr
}

func main() {
	arr := []int{2, 42, 5, 627, 4, 88, 92, 2, 5, 8}

	sortArr := SelectSort(arr)

	fmt.Println(sortArr)
}
