package main

import (
	"fmt"
	"strings"
	"strconv"
)
func find_occurrences(textToSearch string, subText string) string {
	var ans string
	if textToSearch == "" || subText == "" {
		ans = "<No Output>"
		return ans
	}
	textToSearchLower := strings.ToLower(textToSearch)
	subTextLower := strings.ToLower(subText)
	if strings.Contains(textToSearchLower, subTextLower){
		ch := subTextLower[0]
		for i, chS := range textToSearchLower {
			if string(chS) == string(ch) {
				// if the first character is same, compare the rest characters
				end := i+len(subTextLower)
				if end > len(textToSearchLower) {
					end = len(textToSearchLower)
				} 
				if textToSearchLower[i:end] == subTextLower{
					ans += strconv.Itoa(i+1) + ","
					i += len(subTextLower)
				}
			}
		}
		ans = strings.TrimSuffix(ans, ",")
	} else {
		ans = "<No Output>"
	}

	return ans
}
func main() {
	search := "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"
	sub := "Peter "
	fmt.Println(find_occurrences(search, sub))
}