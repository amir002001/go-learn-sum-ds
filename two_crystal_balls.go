package main

import (
	"fmt"
	"math"
)

func TestTwoCrystalBalls() {
	breaks := []bool{false, false, false, false, false, false, true, true, true, true, true, true, true}
	fmt.Println(TwoCrystalBalls(breaks))
}
func TwoCrystalBalls(breaks []bool) int {
	sqrtFloors := int(math.Sqrt(float64(len(breaks))))
	for i := 0; i < sqrtFloors; i++ {
		floorToCheck := (i+1)*sqrtFloors - 1
		if breaks[floorToCheck] == false {
			continue
		}
		for j := sqrtFloors * i; j < sqrtFloors*(i+1); j++ {
			if breaks[j] == true {
				return j
			}
		}

	}
	return -1
}
