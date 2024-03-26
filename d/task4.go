package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome("Abba"))
}
