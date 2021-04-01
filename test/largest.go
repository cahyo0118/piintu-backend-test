package main

import (
	"fmt"
	"strconv"
)

func largest(numbers []int64) int {
	string_concatenating := ""

	for _, n := range numbers {
		string_concatenating += strconv.Itoa(int(n))
	}

	v, _ := strconv.ParseInt(string_concatenating, 10, 32)

	return int(v)
}

func main() {
	var numbers = []int64{1, 10, 100}

	fmt.Println(largest(numbers))
}
