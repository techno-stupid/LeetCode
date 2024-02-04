package main

import "fmt"

func main() {
	nums := []int{3, 2, 4, 7, 11}
	target := 5
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, num := range nums {
		twin := target - num
		if index, found := hashMap[twin]; found {
			return []int{index, i}
		}
		hashMap[num] = i
	}
	return nil
}
