
// 变量是具体数据被内存存储之后内存地址的名称，说白了就是内存中的门牌号。
package main

import (
	f "fmt"
	"unsafe"
)

func main() {
	// 在Golang里我们使用var关键字声明变量
	// 如果是有初始值的变量我们可以省略变量声明中的类型
	// var name int
	// var name =1
	
	// 海象操作符来声明变量，这还重方式可以不使用var关键字。事实上呢，
	// 它更像是一个连贯操作，我们可以既声明又赋值，那么我们这个操作可以算得上是一种赋值表达式
	// 但是需要注意已经声明过的变量就不是再次使用海象操作符来声明了
	// 海象操作符只能用在函数的内部

	// 可以使用&符号查看内存地址
	// 也可以使用unsafe.Sizeof获取内存大小 
	var name int
	name = 666
	f.Println(&name)
	f.Println(unsafe.Sizeof(name))

	// golang 也支持交换赋值操作
	b,a := "bbbbb","aaaaa"

	b,a = a,b

	f.Println(a,b)

}