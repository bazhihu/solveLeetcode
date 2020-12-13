package main

/**
思路：
遍历字符串

*/

func reverseWords(s string) string {
	byteList := []byte(s)

	i, j := 0, 0
	for i < len(s) && j < len(s) {
		if byteList[j] == ' ' {
			byteList = reverse(byteList, i, j)
			i = j + 1
		}
		j++
	}
	byteList = reverse(byteList, i, j)

	byteList = reverse(byteList, 0, len(s))
	return string(byteList)
}

func reverse(b []byte, i, j int) []byte {
	for k := 0; k < (j-i)/2; k++ {
		b[k+i], b[j-1-k] = b[j-1-k], b[k+i]
	}
	return b
}
