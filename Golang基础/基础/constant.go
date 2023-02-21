package main

import (
	f "fmt"
)

func main(){

	// 常量constant 表示不变的值，在程序运行时，不会被代码逻辑修改，比如数字上的圆周率，自然常数e等
	const num int = 1

	// const也可以定义枚举类型
	// 常量声明可以使用 iota 常量生成器初始化，它用于生成一组以相似规则初始化的常量，但是不用每行都写一遍初始化表达式。
	// 在一个 const 声明语句中，在第一个声明的常量所在的行，iota 将会被置为 0，然后在每一个有常量声明的行加一。

	const (
		// 可以在const()中添加一个关键字iota, 每行的iota都会累加1, 第一行的iota默认是0
		a = iota      // iota = 0
		b             // iota = 1
		c             // iota = 2
	)

	f.Println(num)

}