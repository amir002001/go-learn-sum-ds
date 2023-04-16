package main

import "fmt"

func TestLinearSearch() {
	slice := []int{3, 4, 5, 2}
	fmt.Println(LinearSearch(slice, 5))
}

func LinearSearch(hayStack []int, needle int) bool {
	for _, v := range hayStack {
		if v == needle {
			return true
		}
	}
	return false
}
