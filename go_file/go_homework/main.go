package main

import (
	"fmt"

	"github.com/MetaNodeAcademy/study_record/tree/main/go_file/go_homework/palindromicnumber"
	"github.com/MetaNodeAcademy/study_record/tree/main/go_file/go_homework/validparentheses"
)

func main() {
	var num int
	num = 2
	switch num {
	case 1:
		//回文数
		var x int
		fmt.Scan(&x)
		fmt.Println(palindromicnumber.IsPalindrome(x))
	case 2:
		// 有效括号
		var str string
		fmt.Scan(&str)
		fmt.Println(validparentheses.IsValid(str))
	default:
		fmt.Println("无效输入")

	}
	return
}
