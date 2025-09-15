package validparentheses

import "fmt"

func IsValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if _, ok := mapping[char]; ok {
			topElement := ' '
			if len(stack) > 0 {
				topElement = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			if mapping[char] != topElement {
				return false
			}
		} else {
			stack = append(stack, char)
		}
	}
	fmt.Println(stack)
	return len(stack) == 0
}
