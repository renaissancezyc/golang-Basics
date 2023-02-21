package main

import (
	f "fmt"
)

/*	1.defer关键字允许我们推迟到函数返回之前(或任意执行return语句之后)一刻才执行
某个语句或函数。(为什么我们要在返回之后才执行这些语句？因为return语句同样可以包含
一些操作，而不是单纯的返回某个值)

*/


func main(){

	test1()
	/*
	加上defer的输出结果：
		我是test1函数里的第一句话
		我是test1函数里的第三句话
		我是test2的第一句话

	去除defer的输出结果：
		我是test1函数里的第一句话
		我是test2的第一句话
		我是test1函数里的第三句话
	*/

	test3()
	// 输出 5 4 3 2 1 0

}



func test1(){
	f.Println("我是test1函数里的第一句话")
	test2()
	f.Println("我是test1函数里的第三句话")
}

func test2(){
	f.Println("我是test2的第一句话")
}


// 当有多个defer行为被注册时，它们会逆序执行（类似栈，既先进后出，后进先出）
func test3(){
	for i := 0;i<=5;i++{
		defer f.Println(i)
	}
}





