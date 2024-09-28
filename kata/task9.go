package main

func Disemvowel(comment string) string {
	arr := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'y': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true, 'Y': true}
	runes := []rune(comment)
	var result []rune
	for _, char := range runes {
		if !arr[char] { // if !arr[h]
			result = append(result, char)
		}
	}

	return string(result)
}

func main() {
	comment := "Hello, how are you?"
	disemvoweled := Disemvowel(comment)
	println(disemvoweled) // Output: Hll, hw r y?
}
