package main

import "fmt"

func TestBinarySearch() {
	arr := []int{2, 4, 5, 8, 9}
	res := BinarySearch(arr, 1)
	fmt.Println(res)

}

func BinarySearch(hayStack []int, needle int) bool {
	lo := 0
	hi := len(hayStack) - 1
	for lo < hi {
		mid := (lo + hi) / 2
		midVal := hayStack[mid]
		if midVal == needle {
			return true
		} else if needle > midVal {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return false
}
