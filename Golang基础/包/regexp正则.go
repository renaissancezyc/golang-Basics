package main

import (
	f "fmt"
	"regexp"
)

var hhh int = 111111

func main(){
	// 手机号
	phone := "13011111111"
	// 匹配规则
	pat := "^1[1-9]\\d{9}$"

	// 正则调用规则
	reg := regexp.MustCompile(pat)

	
	f.Println(reg.MatchString(phone))
	// 输出 True



	

}