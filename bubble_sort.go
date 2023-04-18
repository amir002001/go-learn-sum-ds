package main

import "fmt"

func TestBubbleSort() {
	slice := []int{2, 5, 34, 12, 6, 7, 12, 34}
	BubbleSort(slice)
	fmt.Println(slice)
}

func BubbleSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				tmp := slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = tmp
			}
		}
	}
}
