package main

func Min(a,b int)int{
	if a < b{
		return a
	}else {
		return b
	}
}
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	res := 0
	resTime := 0
	for left != right{
		resTime = Min(height[left],height[right]) * (right - left)
		if resTime >  res{
			res = resTime
		}
		if height[left] > height[right] {
			right--
		}else {
				left ++
		}
	}
	return res
}