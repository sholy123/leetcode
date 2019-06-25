package main

import (
	"fmt"
	"sort"
)
func isEqual(nums1 []int,nums2 []int) bool{
	sort.Ints(nums1)
	sort.Ints(nums2)
	if len(nums2) != len(nums1){
		return false
	}
	for i := 0; i < len(nums1);i++{
		if nums1[i] != nums2[i]{
			return false
		}
	}
	return true
}
func threeSum(nums []int)[][]int{
	result := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums);i++{
		for j := i+1; j < len(nums);j++{
			for k := j+1; k < len(nums);k++{
				if (nums[i]+nums[j]+nums[k]) == 0{
					res := []int{}
					res = append(res, nums[i])
					res = append(res,nums[j])
					res = append(res,nums[k])
					result = append(result, res)
				}
			}
		}
	}
	return result
}
func test(){
	fmt.Println("hello")
}
func main(){
	test()
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	for a,_ := range res{
		fmt.Println(res[a])
	}
	nums1 := []int{1,2,2}
	nums2 := []int{2,2,1}
	fmt.Println(isEqual(nums1,nums2))
}
