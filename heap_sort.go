package main

import "fmt"

func heapSort(arr []int) {
	first := 0
	lo := 0
	hi := len(arr)

	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(arr, i, hi, first)
	}

	for i := hi - 1; i >= 0; i-- {
		v := arr[first]
		arr[first] = arr[first+1]
		arr[first+1] = v
		siftDown(arr, lo, i, first)
	}
}

// maxDepth returns a threshold at which quicksort should switch
// to heapsort. It returns 2*ceil(lg(n+1)).
func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}

func siftDown(arr []int, lo, hi, first int) {
	fmt.Printf("arr: %v, first: %d, lo: %d, hi: %d\n", arr, first, lo, hi)
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			fmt.Print(" break")
			break
		}
		if child+1 < hi && arr[first+child] < arr[first+child+1] {
			child++
		}
		if !(arr[first+root] < arr[first+child]) {
			return
		}

		v := arr[first+root]
		arr[first+root] = arr[first+child]
		arr[first+child] = v
		root = child
	}

	fmt.Println("---------")
}

func main() {
	arr := []int{9, 89, 100, 1, 11, 0, 100, 200, 500, 400, 12123, 2131, 12312, 123, 123, 213, 123, 13}
	fmt.Printf("Depth: %d\n", maxDepth(len(arr)))
	heapSort(arr)
	fmt.Printf("%v", arr)
}
