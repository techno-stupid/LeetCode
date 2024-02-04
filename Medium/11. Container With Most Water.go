package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(maxArea(nums))
}

func maxArea(height []int) int {
	maximumArea := 0
	length := len(height)
	left, right := 0, length-1
	for left < right {
		minHeight := height[left]
		if height[right] < minHeight {
			minHeight = height[right]
		}
		area := minHeight * (right - left)
		if maximumArea < area {
			maximumArea = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maximumArea
}
