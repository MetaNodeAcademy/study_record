package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MetaNodeAcademy/study_record/tree/main/go_file/go_homework/longestcommonprefix"
	"github.com/MetaNodeAcademy/study_record/tree/main/go_file/go_homework/palindromicnumber"
	"github.com/MetaNodeAcademy/study_record/tree/main/go_file/go_homework/validparentheses"
)

func main() {
	var num int
	num = 3
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
	case 3:
		// 最长公共前缀
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)   // 去掉换行
		strs := strings.Split(line, " ") // 按空格拆分
		fmt.Println(longestcommonprefix.LongestCommonPrefix(strs))
	default:
		fmt.Println("无效输入")

	}
	return
}
