package main

import "fmt"

func shakerSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	var left int = 0
	var right int = len(arr) - 1

	for left <= right {
		for i := right; i > left; i-- {
			if arr[i-1] > arr[i] {
				v := arr[i]
				arr[i] = arr[i-1]
				arr[i-1] = v
			}
		}
		left++
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				v := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = v
			}
		}
		right--
	}
}

func main() {
	arr := []int{5, 6, 7, 1, 2}
	shakerSort(arr)
	fmt.Printf("%v", arr)
}
