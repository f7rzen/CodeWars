package main

import "fmt"

func Gimme(array [3]int) int {
	maxEl := array[0]
	minEl := array[0]
	var midIndex int
	for _, i := range array {
		if i > maxEl {
			maxEl = i
		}
		if i < minEl {
			minEl = i
		}
	}

	for j := 0; j < len(array); j++ {
		if array[j] != maxEl && array[j] != minEl {
			midIndex = j
		}
	}
	return midIndex
}

func main() {
	fmt.Println(Gimme([3]int{5, 10, 14}))
}
