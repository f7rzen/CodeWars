package main

import "fmt"

func isValidInput(input int) bool {
	return 1 <= input && input <= 100
}

func main() {
	var A, B, C, D int
	fmt.Scanln(&A, &B, &C, &D)
	if isValidInput(A) && isValidInput(B) && isValidInput(C) && isValidInput(D) {
		if D > B {
			fmt.Println(A + C*(D-B))
		} else {
			fmt.Println(A)
		}
	} else {
		fmt.Println("Error")
	}
}
