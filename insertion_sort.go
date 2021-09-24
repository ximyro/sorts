package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		x := arr[i]
		j := i
		for j > 0 && (arr[j-1] > x) {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = x
	}
}

func main() {
	arr := []int{5, 7, 4, 2, 1}
	insertionSort(arr)
	fmt.Printf("%v", arr)
}
