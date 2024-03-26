package main

import (
	"fmt"
	"strings"
)

type MyString string

var str MyString

func (s MyString) IsUpperCase() bool {
	return string(s) == strings.ToUpper(string(s))
}

func main() {
	fmt.Println(str.IsUpperCase())
}
