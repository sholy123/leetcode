package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums)  == 0{
		return 0
	}
	length := 0
	for i:=0;i<len(nums);{
		if nums[i] == val{
			i++
		}else{
			nums[length] = nums[i]
			i++
			length++
		}
	}
	return length
}
func divide(dividend int, divisor int) int {
	res := 0
	sign := 1
	if dividend < 0{
		sign = -sign
		dividend = -dividend
	}
	if divisor < 0{
		sign = -sign
		divisor = -divisor
	}
	for dividend >= divisor{
		dividend = dividend - divisor
		res++
	}
	if sign > 0{
		return res
	}else{
		return -res
	}
}
func main(){
	nums := []int{1,2,2,3,4,5}
	res := removeElement(nums,2)
	fmt.Println(res)
	fmt.Println(nums[:res])
	fmt.Println(divide(7,-3))
}
