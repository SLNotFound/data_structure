package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrome(x int) {
	// 将数字转换成字符串
	s := strconv.Itoa(x)
	// 将字符串转成字符数组
	r := strings.Split(s, "")

	// 翻转字符数组，从后往前排列
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	fmt.Println(strings.Join(r, "") == s)
}

func main() {
	isPalindrome(121)
}
