package main

import (
	f "fmt"
)


func main(){

	f.Println("请输入一个整数")
	var x int 

	// 获取输入  Scanln输入
	f.Scanln(&x)  //这里是通过指针来接入给变量x的意思

	f.Println(x)

	var bin = 1 >> 2

	f.Println(bin)



}