package palindromicnumber

import "fmt"

// 回文数
func IsPalindrome() {
	var str string
	var num int
	var strlen int
	for {
		fmt.Scan(&str)
		strlen = len(str)
		num = strlen
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num/2 + 1
		}
		var char rune
		for i, v := range str {
			is := false
			index := strlen
			for _, vv := range str {
				index--
				if index < num {
					is = false
					break
				}
				if index == i {
					char = vv
					is = true
					break
				}
			}
			if is {
				if v != char {
					fmt.Print("false")
					break
				} else {
					if i == num-1 {
						fmt.Print("true")
						break
					}
				}
			}
		}
	}
	return
}
