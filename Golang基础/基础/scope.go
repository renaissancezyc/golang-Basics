package main

import (f "fmt")


// 在函数或者块之外定义的变量我们称之为全局变量，这些变量在程序的整个生命周期中国，都可以使用
var myvarable3 int = 6


func main(){
	// 变量的作用域可以理解为可访问变量的程序的某个范围，比如说啊，在类里，方法里
	// 循环里面定义的变量，那么在函数或代码库尔中声明的变量，我们称之为虎局部变量，
	// 这些变量不能在函数或块之外的的地方访问，所以这些变量我们可以称之为块变量。

	var myvarable1,myvarable2 int =  69,145

	f.Println(myvarable1,myvarable2,myvarable3)

}