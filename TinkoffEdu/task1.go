package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if -32000 <= x && x <= 32000 && -32000 <= y && y <= 32000 {
		fmt.Println(x + y)
	} else {
		fmt.Println("error")
	}
}
