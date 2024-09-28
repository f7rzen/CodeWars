package main

import "fmt"

var romanNumerals = []struct {
	Value  int
	Symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func Solution(number int) string {
	result := ""
	for _, roman := range romanNumerals {
		for number >= roman.Value {
			result += roman.Symbol
			number -= roman.Value
		}
	}
	return result
}

func main() {
	fmt.Println(Solution(1666))
}
