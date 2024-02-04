package main

import "strconv"

func main() {
	x := -91
	println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	rev, _ := strconv.Atoi(reverse(strconv.Itoa(x)))
	if rev == x {
		return true
	}
	return false
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
