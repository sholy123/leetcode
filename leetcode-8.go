package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func myAtoi(str string) int {
	var res = 0
	var flag = false
	for i := 0;i < len(str);i++{
		if str[i] == ' '{
			continue
		}else if str[i] == '-'{
			flag = true
		}else if str[i] < '0' || str[i] > '9'{
			break
		}else {
			ta := str[i] - 48
			res = res * 10 + int(ta)
		}
	}
	if flag{
		return -res
	}else{
		return res
	}
}

func main(){
	str := "7979179gah"
	intTest := 1334155
	res, _ := strconv.Atoi(str)
	str2 := strconv.Itoa(intTest)

	fmt.Println(reflect.TypeOf(str2))
	fmt.Println(res)
	fmt.Println(reflect.TypeOf(myAtoi(str)))
}