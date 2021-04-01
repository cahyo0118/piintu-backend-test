package main

import (
	"fmt"
	"strconv"
)

func largest(numbers []int64) string {
	string_concatenating := ""

	for _, n := range numbers {
		string_concatenating += strconv.Itoa(int(n))
	}

	return string_concatenating
}

func main() {
	var numbers = []int64{1, 10, 100}

	fmt.Println(largest(numbers))
}
