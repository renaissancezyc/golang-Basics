package main

import (
	f "fmt"
)

func main(){
	
	// 数学运算符
	//  + - * / %(求余)
	a := 1
	b := 2
	c := false
	f.Println(a+b)
	f.Println(a-b)
	f.Println(a*b)
	f.Println(a/b)
	f.Println(a%b)

// 比较运算符
	// 等于 不等于 大于 小于 大于等于 小于等于

if (a==b){
	f.Println("等于")
}

if (a!=b){
	f.Println("不等于")
}

if (a>b){
	f.Println("大于")
}

if (a<b){
	f.Println("小于")
}

if (a>=b){
	f.Println("大于等于")
}

if (a<b){
	f.Println("小于等于")
}


// 逻辑运算符
// 逻辑与运算符 逻辑或运算符 逻辑非运算符

if a==1 && b==2{
	f.Println("逻辑与")
}

if a==1 || b==1{
	f.Println("逻辑或")
}

if !c{
	f.Println("逻辑非")
}


// 位运算
	f.Println(a & b) //位于  可以判断一个结点是否有一个权限
	f.Println(a | b) //位或   可以给一个结点权限
	f.Println(a ^ b) //位异或  可以删除一个结点的权限
	f.Println(a << 1) //左位移  可以将某一条数据整体向左移动
	f.Println(a >>1) //右位移   可以将某一条数据整体向右移动


// 捕获用户在终端输入的数据

var  x int
f.Println("请输入一个整数")
f.Scanln(&x) //读取键盘的输入，通过操作地址，赋值给x
f.Println(x)


}