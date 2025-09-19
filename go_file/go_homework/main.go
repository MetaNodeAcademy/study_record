package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/MetaNodeAcademy/study_record/tree/main/go_file/go_homework/longestcommonprefix"
	"github.com/MetaNodeAcademy/study_record/tree/main/go_file/go_homework/mergeintervals"
	"github.com/MetaNodeAcademy/study_record/tree/main/go_file/go_homework/palindromicnumber"
	"github.com/MetaNodeAcademy/study_record/tree/main/go_file/go_homework/plusone"
	"github.com/MetaNodeAcademy/study_record/tree/main/go_file/go_homework/removeduplicates"
	"github.com/MetaNodeAcademy/study_record/tree/main/go_file/go_homework/thenumberthatappearsonlyonce"
	"github.com/MetaNodeAcademy/study_record/tree/main/go_file/go_homework/twonumberssum"
	"github.com/MetaNodeAcademy/study_record/tree/main/go_file/go_homework/validparentheses"
)

func main() {
	num := 8
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
	case 4:
		// 只出现一次的数字
		var nums []int
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		strs := strings.Split(line, " ")
		for _, v := range strs {
			n, _ := strconv.Atoi(v)
			nums = append(nums, n)
		}
		fmt.Println("只出现一次的数字是：", thenumberthatappearsonlyonce.SingleNumber(nums))
	case 5:
		// 加一
		var digits []int
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		strs := strings.Split(line, ",")
		for _, v := range strs {
			n, _ := strconv.Atoi(v)
			digits = append(digits, n)
		}
		result := plusone.PlusOne(digits)
		fmt.Println("加一后的结果是：", result)
	case 6:
		// 删除有序数组中的重复项
		var digits []int
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		strs := strings.Split(line, ",")
		for _, v := range strs {
			n, _ := strconv.Atoi(v)
			digits = append(digits, n)
		}
		fmt.Println("删除有序数组中的重复项", removeduplicates.RemoveDuplicates(digits))
	case 7:
		// 合并区间
		var intervals [][]int
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		_ = json.Unmarshal([]byte(line), &intervals)
		fmt.Println("合并区间", mergeintervals.MergeIntervals(intervals))
	case 8:
		nums := []int{}
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		_ = json.Unmarshal([]byte(line), &nums)
		var target int
		fmt.Scan(&target)
		fmt.Println("两数之和", twonumberssum.TwoSum(nums, target))
	default:
		fmt.Println("无效输入")

	}
}
