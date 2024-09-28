package main

import (
	"fmt"
	"strconv"
)

func StringToNumber(str string) int {
	x, _ := strconv.Atoi(str)
	return x
}

func main() {
	fmt.Println(StringToNumber("1234"))
}
