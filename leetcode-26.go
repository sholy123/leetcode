package main

import "fmt"

func removeDuplicates(nums []int) int {
	length := 0
	if len(nums) <= 1{
		return len(nums)
	}
	for i:=0; i < len(nums);i++{
		if nums[i] != nums[length]{
			length++
			nums[length] = nums[i]
		}else{
				continue
		}
	}
	return length+1
}
func main(){
	nums := []int{1,1,2,3,4,5,5}
	re := removeDuplicates(nums)
	fmt.Println(nums[:re])
}