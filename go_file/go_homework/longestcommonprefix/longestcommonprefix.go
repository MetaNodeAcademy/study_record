package longestcommonprefix

// 最长公共前缀
func LongestCommonPrefix(strs []string) string {
	str := []rune{}
	for i := 0; ; i++ {
		for num, s := range strs {
			if i >= len(s) {
				if num != 0 {
					return string(str[:len(str)-1])
				} else {
					return string(str)
				}
			}
			if num == 0 {
				str = append(str, rune(s[i]))
			} else {
				if rune(s[i]) != str[i] {
					return string(str[:i])
				}
			}
		}
	}
	return string(str)
}
