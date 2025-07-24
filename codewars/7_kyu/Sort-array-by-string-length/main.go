package main

import "fmt"

func mergeSort(arr []string) []string {
	// это когда один элемент
	if len(arr) < 2 {
		return arr
	}

	// делим на 2 массива
	mid := len(arr) / 2
	// делим хуй пойми как ну и похуй
	left, right := mergeSort(arr[:mid]), mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []string) []string {
	var merged []string
	for len(left) > 0 && len(right) > 0 {
		if len(left[0]) <= len(right[0]) {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	merged = append(merged, left...)
	merged = append(merged, right...)
	return merged
}

func SortByLength(arr []string) []string {

	return mergeSort(arr)
}

func main() {
	data := []string{"", "eloquent", "bees", "dictionary"}
	fmt.Println(SortByLength(data))
}
