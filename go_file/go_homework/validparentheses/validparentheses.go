package validparentheses

// 有效的括号
func IsValid(s string) bool {

	strlen := len(s)
	if strlen%2 != 0 {
		return false
	}
	isvalid := false
	var str string
	for i := 0; ; i++ {
		if i >= strlen {
			isvalid = true
			break
		}
		if s[i] != '(' && s[i] != '{' && s[i] != '[' {
			isvalid = false
			break
		}
		str = ""
		if s[i] == '(' {
			str = string(s[i])
			j := i
			for {
				i++
				if i >= strlen {
					return false
				}
				str = str + string(s[i])
				if s[i] == ')' {
					if i-j == 1 {
						isvalid = true
						break
					}
					str = str[1:]
					str = str[:len(str)-1]
					if IsValid(str) {
						isvalid = true
						break
					} else {
						isvalid = false
					}

				}
			}
		}
		if s[i] == '[' {
			str = string(s[i])
			j := i
			for {
				i++
				if i >= strlen {
					return false
				}
				str = str + string(s[i])
				if s[i] == ']' {
					if i-j == 1 {
						isvalid = true
						break
					}
					str = str[1:]
					str = str[:len(str)-1]
					if IsValid(str) {
						isvalid = true
						break
					} else {
						isvalid = false
					}

				}
			}
		}
		if s[i] == '{' {
			str = string(s[i])
			j := i
			for {
				i++
				if i >= strlen {
					return false
				}
				str = str + string(s[i])
				if s[i] == '}' {
					if i-j == 1 {
						isvalid = true
						break
					}
					str = str[1:]
					str = str[:len(str)-1]
					if IsValid(str) {
						isvalid = true
						break
					} else {
						isvalid = false
					}

				}
			}
		}
	}
	return isvalid
}
