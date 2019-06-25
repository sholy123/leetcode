package main

import "fmt"

func MaxSub(str1 string, str2 string)int {
	res := 0
	//i := len(str1)
	//j := len(str2)
	dp := [][]int{}
	for i := 0; i < len(str1);i++{
		for j := 0; j < len(str2);j++{
			dp[i][j] = 0
			fmt.Println(i,j)
		}
	}
	for i := 1; i <= len(str1);i++{
		for j := 1; j <= len(str2);j++{
			if str1[i-1] == str2[j-1]{
				dp[i][j] = dp[i-1][j-1]+1
				res += 1
			}else{
				dp[i][j] = 0
			}
		}
	}
	return res
}
func main() {
	str1 := "hello"
	str2 := "hell"
	fmt.Println(MaxSub(str2,str1))
}
