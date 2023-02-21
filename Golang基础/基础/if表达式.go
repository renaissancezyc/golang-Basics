package main

import f "fmt"



func test1(a int) int {

	
	// 一个简单的if语句
	if a>0{
		return a
	}else{
		return 0000000
	}

	
}


func test2(a int) int {
	// if语句可以在条件表达式前执行一个简单的语句
	// 可以看到我在if表达式前声明了一个num变量
	// 该变量的作用域仅在if语句内
	if num:=a+10 ; num>20{
		return num
	} else if num<20{
		return 20
	}else{
		return 0000000
	}
	
}


func main(){

	f.Println(test1(-1))
	f.Println(test2(4))


}