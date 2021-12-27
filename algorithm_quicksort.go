package main

import "fmt"

func quickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
	if left >= right {
		return []int{}
	}

	l, r, pivot := left, right, arr[left]
	// {92, 42, 5, 627, 4, 88, 95, 2, 5, 8}
	for l < r {
		for l < r && pivot <= arr[r] {
			r -= 1
		}

		if l < r && pivot >= arr[r] {
			arr[l] = arr[r]
		}
		for l < r && pivot >= arr[l] {
			l += 1
		}
		if l < r && pivot <= arr[l] {
			arr[r] = arr[l]
		}
		if l >= r {
			arr[l] = pivot
		}
	}

	fmt.Println(arr)
	_quickSort(arr, left, l)
	_quickSort(arr, l+1, right)
	return arr
}
func main() {
	arr := []int{92, 42, 5, 627, 4, 88, 95, 2, 5, 8}

	sortArr := quickSort(arr)

	fmt.Println(sortArr)
}
