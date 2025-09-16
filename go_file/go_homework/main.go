package main

import (
	"fmt"

	"github.com/MetaNodeAcademy/study_record/tree/main/go_file/go_homework/palindromicnumber"
	"github.com/MetaNodeAcademy/study_record/tree/main/go_file/go_homework/validparentheses"
)

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(palindromicnumber.IsPalindrome(x))
	validparentheses.IsValid()
	return
}
