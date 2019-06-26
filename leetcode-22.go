package main

import (
	"fmt"
)

func generateParenthesis(n int)[]string{
	res := []string{}
	return res
}
func generalALL(current string,pos int, result []string){
	if (pos == len(current)){
		if isValid(current){
			result = append(result,current)
		}
	}else{
		string(current[pos]) = "("
		generalALL(current,pos+1,result)
		string(current[pos]) = ")"
		generalALL(current,pos+1,result)
	}
}
func isValid(str string)bool{
	balance := 0
	for i:=0; i < len(str);i++{

		if string(str[i]) == "("{
			balance ++
		}else{
				balance --
		}
	}
	if balance == 0{
		return true
	}else{
		return false
	}
}
func main() {
	str := "((())())"
	fmt.Println(isValid(str))
}