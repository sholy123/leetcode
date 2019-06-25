package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	/* location[s[i]] == j 表示：s中第i个字符串，上次出现在s的j位置，所以，在s[j+1:i]中没有s[i]
	   location[s[i]] == -1 表示：s[i] 在s中第一次出现*/
	var length int
	location := [256]int{}
	for i := range location{
		location[i] = -1
	}
	length = 0
	left := 0
	for i := 0; i < len(s);i++{
		if location[s[i]] >= left {
			left = location[s[i]] + 1
		}else if i+1-left > length{
			length = i+1 - left
		}
		location[s[i]] = i
	}
	return length
}

func main(){
	s := "pppppppp"
	fmt.Println(lengthOfLongestSubstring(s))
}