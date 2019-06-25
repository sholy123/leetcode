package main

import (
	"fmt"
)

func fib(n int) []int{
	mero := make([]int,n+1)
	mero[0] = 0
	mero[1] = 1
	for i:=2; i <=n; i++{
		mero[i] = mero[i-1] + mero[i-2]
	}
	return mero
}
func max(m,n int) int{
	if m > n{
		return m
	}else{
		return n
	}
}
func cut(p []int,n int) int{
	if n == 0{
		return 0
	}
	q := 0
	for i :=1 ;i <= n;i++{
		q = max(q,p[i-1]+cut(p,n-i))
	}
	return q
}
func reverse(n int)int{
	res := 0
	a := 0
	for n != 0{
		a = n % 10
		n = n / 10
		res = res * 10 + a
	}
	return res
}
func main(){
	fmt.Println(fib(6))
	p := []int{1,5,8,9,10,17,17,20,24,30}
	fmt.Println(cut(p,5))
	fmt.Println(reverse(-555))
}
