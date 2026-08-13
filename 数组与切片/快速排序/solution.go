package main

import "fmt"

func quicksort(arr []int, left int, right int) {
	if left >= right {
		return
	}
	i := left
	j := right
	pivot := arr[left]
	for i < j {
		for i < j && arr[j] >= pivot {
			j--
		}
		for i < j && arr[i] <= pivot {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[left], arr[i] = arr[i], arr[left]
	quicksort(arr, left, i-1)
	quicksort(arr, i+1, right)
}

func main() {
	arr := []int{1, 3, 2, 5, 9, 7}
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
