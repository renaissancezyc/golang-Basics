package main

import f "fmt"


func main(){

	// switch 是编写一连串 if-else 语句的简便方法。
	// 它允许第一个值等于条件表达式的case语句
	// 每一个case都相当于一个elif
	// Golang自动给这些语句中每个case后面添加了break语句(跳出循环)
	// switch的case语句从上到下顺次执行,直到匹配成功时停止，如果没有满足条件的case，也可以设置默认值default。
	// 使用switch可以将一长串的 if-then-else 变得便于阅读
	switch num:= "你好";num{
		
		case "你好":
			f.Println("谢谢")
		
		case "吃饭了么":
			f.Println("没吃")

		case "Golang","Go":
			f.Println("Go语言")
		
		default:
			f.Println("请输入我识别的话")
	}

	//结果: 谢谢 

	

}