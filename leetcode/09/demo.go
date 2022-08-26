package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrome(x int) {
	str := strconv.Itoa(x)
	split := strings.Split(str, "")
	temp := split[:len(split)/2]
	strings.Join(temp, "")
	fmt.Println(strings.Join(temp, ""))
}

func test(x int) {
	s := strconv.Itoa(x)
	r := strings.Split(s, "")
	//r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	//return string(r) == s

	fmt.Println(strings.Join(r, "") == s)

	//count := len(split)
	//
	//for i := len(split); i < len(split)/2+1; i-- {
	//	if x/int(math.Pow10(i)) == x%int(math.Pow10(count-i+1)) {
	//		fmt.Println("123")
	//	}
	//	x = x - int(math.Pow10(i))
	//}
}

func main() {
	test(12321)
}
