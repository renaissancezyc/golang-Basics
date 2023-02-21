package main

import f "fmt"

// 函数练习2

/*
	init函数这是一类非常特殊的函数，它不能被认为调用，而是在每个包完成初始化后自动执行，
	并且执行优先级比 main 函数高。

	每个源文件都只能包含一个init函数。初始化总是以单线程执行，并且按照包的以来关系顺序执行。
*/


func init(){
	f.Println("我是init函数")
}

func main(){
	f.Println("我是main函数")
}