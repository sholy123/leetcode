package main

import (
	"fmt"
)

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
func permute(nums []int) [][]int {
	length := len(nums) - 1
	var res [][]int
	var n []int
	dfs1(nums, 0, length, &n)
	for m := 0; m < len(n); m = m + len(nums) {
		res = append(res, n[m:m+len(nums)])
	}
	return res
}
func dfs1(nums []int, cur, end int, res *[]int) {
	if cur == end {
		*res = append(*res, nums...)
		return
	} else {
		for i := cur; i <= end; i++ {
			swap(nums, cur, i)
			dfs1(nums, cur+1, end, res)
			swap(nums, cur, i)
		}
	}
}
func main(){
	nums := []int{1,2,3}
	fmt.Println(permute(nums))
}