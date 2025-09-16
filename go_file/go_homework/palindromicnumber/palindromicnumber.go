package palindromicnumber

import "math"

// 回文数
func IsPalindrome(x int) bool {
	//===========3============================//
	if x < 0 {
		return false
	}
	ispalindrome := false
	charmap := map[int]int{}
	input := x
	for i := 1; ; i++ {
		charmap[i] = input % int(math.Pow10(1))
		input = input / int(math.Pow10(1))
		if (input < 10) && (input > -10) {
			charmap[i+1] = input
			break
		}
	}
	num := len(charmap)
	for i := 1; ; i++ {
		if i == num-i+1 {
			ispalindrome = true
			break
		}
		if charmap[i] != charmap[num-i+1] {
			ispalindrome = false
			break
		}
	}
	return ispalindrome
	//===========2============================//
	// var input int = x
	// ispalindrome := false
	// // if x < 0 {
	// // 	input = -x
	// // }
	// inputstr := strconv.Itoa(input)
	// num := len(inputstr)
	// num--
	// for i, v := range inputstr {
	// 	if i == num-i {
	// 		ispalindrome = true
	// 		break
	// 	}
	// 	if v != rune(inputstr[num-i]) {
	// 		ispalindrome = false
	// 		break
	// 	}
	// }
	// return ispalindrome
	//===========1============================//
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
