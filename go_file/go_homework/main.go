package main

import (
	"fmt"
	"validparentheses"
)

func main() {
	testCases := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
	}

	for _, testCase := range testCases {
		result := validparentheses.IsValid(testCase)
		fmt.Printf("Input: %s, Output: %v\n", testCase, result)
	}

	validparentheses.IsValid("()[]{}")
	return
}
