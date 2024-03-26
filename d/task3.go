package main

import "fmt"

func Grow(arr []int) int {
	result := 1
	for i := range arr {
		result *= arr[i]
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 5, 6}
	fmt.Println(Grow(arr))
}
