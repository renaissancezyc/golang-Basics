package main

import f "fmt"

func main() {

	test1()

}

/* 使用defer语句实现代码追踪:
一个比较实用的功能实现代码追踪的方案就是在进入和离开某个函数时打印相关信息。
*/

// 追踪函数
func trace(s string) string {
	f.Println("entering:", s)
	return s
}

func untrace(s string) {
	f.Println("leaving:", s)
}

func test1() {
	// 因为传参不受限与defer关键字，所以第一步执行的是trace函数，然后到最后将返回值
	// 给到untrace才执行untrace
	defer untrace(trace("test1"))
	f.Println("我现在在test1函数内")
}
