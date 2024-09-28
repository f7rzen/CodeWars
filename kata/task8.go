package main

import "fmt"

func GetCount(str string) (count int) {
	slize := []string{"a", "e", "i", "o", "u"}
	for _, value := range slize {
		if value == str {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(GetCount("abracadabra"))
}
