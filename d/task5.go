package main

import (
	"fmt"
	"sort"
)

func TwoSort(arr []string) string {
	sort.Strings(arr)
	var sortString string

	for i, char := range arr[0] {
		// Добавляем букву к строке
		sortString += string(char)

		// Если это не последняя буква, добавляем разделитель "***"
		if i < len(arr[0])-1 {
			sortString += "***"
		}
	}

	return sortString
}

func main() {
	s := []string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}
	fmt.Println(TwoSort(s))
	for i, i2 := range s {
		fmt.Println(i, i2)
	}
}
