package main

import "fmt"

func bubleSort(arr []int) {
	for i := 0; i+1 < len(arr); i++ {
		for j := 0; j+1 < len(arr)-i; j++ {
			if arr[j+1] < arr[j] {
				v := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = v
			}
		}
	}
}

func main() {
	arr := []int{5, 8, 3, 2, 1, 1}
	bubleSort(arr)
	fmt.Printf("%v", arr)
}
