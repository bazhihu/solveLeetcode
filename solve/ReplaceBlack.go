package solve

// 替换字符串空格

import (
	"fmt"
)

func ReplaceBlack(str []rune) {
	strLen := len(str)
	if strLen <= 0 {
		return
	}

	var spaceLen int
	for i := 0; i < strLen; i++ {
		if string(str[i]) == " " {
			spaceLen++
		}
	}

	if spaceLen <= 0 {
		return
	}

	strNewLen := strLen + spaceLen*3

	j := strNewLen - 1

	var a rune = '0'
	var b rune = '2'
	var c rune = '%'

	newStr := make([]rune, 50)
	for i := strLen - 1; i >= 0; i-- {

		if string(str[i]) == " " {
			newStr[j] = a
			newStr[j-1] = b
			newStr[j-2] = c
			j = j - 3
		} else {
			newStr[j] = str[i]
			j--
		}
	}

	fmt.Println(string(newStr))
}

func main() {
	str := []rune("sdada asdad adasd asda")
	ReplaceBlack(str)
}
