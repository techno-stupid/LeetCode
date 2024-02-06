package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(str))
}

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	anagramMap := make(map[string][]string)
	for _, str := range strs {
		sortedStr := sortString(str)
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}
	for _, anagramGroup := range anagramMap {
		result = append(result, anagramGroup)
	}
	return result
}

func sortString(input string) string {
	characters := strings.Split(input, "")
	sort.Strings(characters)
	sortedString := strings.Join(characters, "")
	return sortedString
}
