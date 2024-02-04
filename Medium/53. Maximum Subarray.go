package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, -1, -2, 2, 1, -2, 1, 4, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	maxSum, currentSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum += nums[i]
		if currentSum < nums[i] {
			currentSum = nums[i]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}
