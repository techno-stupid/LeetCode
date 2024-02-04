package main

import (
	"fmt"
)

func main() {
	nums := ""
	fmt.Println(letterCombinations(nums))
}

var Result []string
var DigitMap = map[string]string{
	"2": "abc", "3": "def", "4": "ghi", "5": "jkl",
	"6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz",
}

func letterCombinations(digits string) []string {
	Result = []string{}
	if len(digits) != 0 {
		backTrack("", digits)
	}
	return Result
}

func backTrack(combination string, remainingDigits string) {
	if len(remainingDigits) == 0 {
		Result = append(Result, combination)
		return
	}
	digit := string(remainingDigits[0])
	letters := DigitMap[digit]
	for _, latter := range letters {
		backTrack(combination+string(latter), remainingDigits[1:])
	}
}
