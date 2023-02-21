package main

import (
	f "fmt"
	"os"
)

/*
	Go语言的函数经常使用两个返回值来表示执行是否成功，返回某个值以及true表示成功；
	返回零值(或nil)和false表示失败。当不使用true和false时，也可以使用一个error类型
	的变量来代替作为第二个返回值，函数执行成功的话，error值为nil，否则就会包含响应的
	错误信息
*/



func main(){

	data,err := os.Open("111.go")
	if err != nil{
		f.Println("检测到了错误；",err)
	}

	f.Println(data)

}