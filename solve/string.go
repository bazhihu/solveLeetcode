package main

/**
 * isUnique 判断字符串中字符是否唯一
 *
 * BM字符串算法
 *
 *
 */

import (
	"fmt"
)

func isUnique(astr string) bool {
	if len(astr) <= 0 {
		return true
	}
	w := 0
	b := int(astr[0])
	for k, _ := range astr {
		w1 := int(astr[k]) ^ w
		fmt.Println(w1)
		if int(astr[k]) < b {
			b = int(astr[k])
		}
		w = w1
		if w < b {
			return false
		}
	}
	return true
}

// 返回匹配字符串的下标
// pri 主串 被匹配的大串
// mod 模式串 需要匹配的子串
func bm(pri, mod string) int {
	if len(pri) <= 0 || len(mod) <= 0 {
		return -1
	}

	priLen := len(pri)
	modLen := len(mod)

	// 坏字符表
	var badTable = func(modLen int) []int {
		var bc = make([]int, 255)

		for i := 0; i < 255; i++ {
			bc[i] = -1
		}

		for i := 0; i < modLen; i++ {
			bc[int(mod[i])] = i
		}
		return bc
	}
	var badT = badTable(modLen)

	// 好后缀
	// 前缀bool表
	_ = func(mod string) (suffix []int, prefix []bool) {
		suffix = make([]int, 255)
		prefix = make([]bool, 255)

		for i := 0; i < 255; i++ {
			suffix[i] = -1
			prefix[i] = false
		}

		for i := 0; i < modLen-1; i++ {
			j := i
			k := 0                                    // 公共后缀子串长度
			for j >= 0 && mod[j] == mod[modLen-1-k] { // 与mod[0, modLen-1] 求公共后缀子串
				j--
				k++
				suffix[k] = j + 1 //j+1表示公共后缀子串在mod[0, i]中的起始下标
			}
			if j == -1 {
				prefix[k] = true //如果公共后缀子串也是模式串的前缀子串
			}
		}

		return
	}

	var i int = 0            // 表示主串与模式串对齐的第一个字符串
	for i <= priLen-modLen { // 模式串 移动的长度
		var j int
		for j = modLen - 1; j >= 0; j-- { // 模式串从后往前匹配
			if pri[i+j] != mod[j] { // 坏字符对应模式串中的下标是
				break
			}
		}
		if j < 0 {
			return i // 匹配成功，返回从左到右 主串与模式串
		}

		i = i + j - badT[int(pri[i+j])]
	}

	return -1
}

func main() {
	//a := isUnique("aa")

	//kk := bm("fbajjbabajksks", "baba")
	//fmt.Println(kk)
}
