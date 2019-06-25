package main

import "fmt"

func traverse(s string, numRows int) string{
	if numRows == 1{
		return s
	}
	res := ""
	j := 0
	jiange := (numRows - 1) * 2
	var i int
	var flag bool
	for i = 1; i <= numRows;i++{
		//for j := 0; j <= len(s);j= j + jiange{
		//	res = res + string(s[j])
		//}
		flag = true
		for j < len(s){
			res = res +string(s[j])
			if jiange == ((numRows-1) * 2){
				j = j + jiange
			}else if flag {
				j = j + jiange
				flag = false
			}else{
				j = j + ((numRows-1)*2) - jiange
				flag = true
			}
		}
		j = i
		jiange = jiange - 2
		if i == numRows-1 {
			jiange = (numRows - 1) * 2
		}
	}
	return  res
}

func main(){
	str := "LEETCODEISHIRING"
	res := traverse(str,4)
	fmt.Println(res)
}
