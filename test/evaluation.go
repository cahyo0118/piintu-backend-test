package main

import (
	"fmt"
	"strconv"
	"strings"
)

func evaluation(input string) int64 {

	input = strings.ReplaceAll(input, " ", "")
	input = strings.ToLower(input)

	// nurcahyo0118 -> more effective if using regex
	var characters = strings.Split(input, "")
	var formula = []string{}
	tempValue := ""

	for _, c := range characters {

		if _, err := strconv.Atoi(c); err == nil {
			tempValue += c
		} else if c == "+" {
			formula = append(formula, tempValue)
			tempValue = ""
			formula = append(formula, c)
		} else if c == "-" {
			formula = append(formula, tempValue)
			tempValue = ""
			formula = append(formula, c)
		}

	}

	// Append last value
	formula = append(formula, tempValue)

	var finalValue int64 = 0

	if len(formula) > 0 {

		firstNumber, _ := strconv.ParseInt(formula[0], 10, 32)
		finalValue += firstNumber

		for i, f := range formula {

			if f == "+" {
				nextNumber, _ := strconv.ParseInt(formula[i+1], 10, 32)
				finalValue += nextNumber
			} else if f == "-" {
				nextNumber, _ := strconv.ParseInt(formula[i+1], 10, 32)
				finalValue -= nextNumber
			}
		}
	}

	return finalValue
}

func main() {
	fmt.Println(evaluation("1+1+1-1"))
}
