package main

import "fmt"

func Parse(data string) []int {
	value := 0
	n := 0
	for _, char := range data {
		if char == '0' {
			n++
		}
	}
	array := make([]int, n)
	for _, char := range data {
		if 'i' == char {
			value++
		} else if 'd' == char {
			value--
		} else if 's' == char {
			value = value * value
		} else if 'o' == char {
			array = append(array, value)
		}
	}
	return array
}

func main() {
	fmt.Println(Parse("iiisdoso"))
}
