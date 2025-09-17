package plusone

import (
	"math"
	"strconv"
)

func PlusOne(digits []int) []int {
	num := len(digits)
	number := 0
	for i := 0; i < num; i++ {
		number += digits[i] * int(math.Pow10(num-i-1))
	}
	number++
	numstr := strconv.Itoa(number)
	num = len(numstr)
	result := []int{}
	for i := 0; i < num; i++ {
		n, _ := strconv.Atoi(string(numstr[i]))
		result = append(result, n)
	}

	return result
}
