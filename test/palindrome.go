package main

import (
	"fmt"
	"strings"
)

func isPalindrome(input string) bool {

	input = strings.ReplaceAll(input, " ", "")
	input = strings.ToLower(input)

	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("anNa"))
	fmt.Println(isPalindrome("anNNa"))
	fmt.Println(isPalindrome("not a palindrome"))
}
