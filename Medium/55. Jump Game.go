package main

import (
	"fmt"
)

func main() {
	x := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(x))
}

func canJump(nums []int) bool {
	lastPosition := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= lastPosition {
			lastPosition = i
		}
	}
	return lastPosition == 0
}
