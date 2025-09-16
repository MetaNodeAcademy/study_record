package palindromicnumber

import "math"

// 回文数
func IsPalindrome(x int) bool {
	var input int
	ispalindrome := false
	if x < 0 {
		input = -x
	}
	num := 0
	for i := 0; ; i++ {
		if 0 < input%int(math.Pow10(i)) && int(math.Pow10(i)) < 10 {
			num = i
			break
		}
	}
	for i := 0; ; i++ {
		if i == num {
			ispalindrome = true
			break
		}
		if input%int(math.Pow10(i)) != input%int(math.Pow10(num)) {
			ispalindrome = false
			break
		}
	}

	return ispalindrome
	// var str string
	// var num int
	// var strlen int
	// for {
	// 	fmt.Print("请输入一个字符串：")
	// 	fmt.Scan(&str)
	// 	str = strings.TrimLeft(str, "+-")
	// 	strlen = len(str)
	// 	num = strlen
	// 	if num%2 == 0 {
	// 		num = num / 2
	// 	} else {
	// 		num = num/2 + 1
	// 	}
	// 	var char rune
	// 	for i, v := range str {
	// 		is := false
	// 		index := strlen
	// 		for _, vv := range str {
	// 			index--
	// 			if index == i {
	// 				char = vv
	// 				is = true
	// 				break
	// 			}
	// 		}
	// 		if is {
	// 			if v != char {
	// 				fmt.Print("false")
	// 				break
	// 			} else {
	// 				if i == num-1 {
	// 					fmt.Print("true")
	// 					break
	// 				}
	// 			}
	// 		}
	// 	}
	// }
}
