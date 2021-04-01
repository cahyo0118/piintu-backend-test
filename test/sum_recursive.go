package main

import (
	"fmt"
	"math"
)

var nodes []int = []int{}

func addNode(value int) {

	nodes = append(nodes, value)

	if value > 1 {
		addNode(int(float64(value) / 2))

		addNode(int(math.Sqrt(float64(value))))
	}

}

func sum_recursive(number int) int {

	nodes = append(nodes, number)

	if number > 1 {
		addNode(int(float64(number) / 2))

		addNode(int(math.Sqrt(float64(number))))
	}

	result := 0
	for _, v := range nodes {
		result += v
	}
	return result
}

func main() {

	fmt.Println(sum_recursive(7))
	fmt.Println(nodes)
	nodes = []int{}

	fmt.Println(sum_recursive(8))
	fmt.Println(nodes)
	nodes = []int{}

}
