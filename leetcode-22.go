package main

import "fmt"
func generate(n int)[]string{
	res := []string{}
	gen(&res,"",0,n,0)
	return res
}
func gen(res *[]string,temp string,left,n,right int){
	if right == n{
		*res = append(*res,temp)
		return
	}
	if left < n{
		gen(res,temp+"(",left+1, n, right)
	}
	if right < left{
		gen(res, temp + ")",left, n, right + 1)
	}
}
func main() {
	n := 3
	res := generate(n)
	fmt.Println(res)
}