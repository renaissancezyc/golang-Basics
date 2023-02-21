package main   
// golang是以包作为管理单位 每个Golang源文件必须声明它所属的包
// 所以我们会看到每个GO源文件的开头都是一个package声明
// main包是Golang程序的入口包,一个golang语言程序必须包含一个main包
// 那么如果一个程序没有main包 那么编译时就会出错,无法生成可执行文件

// import ."fmt" 这样写可以省略包的前缀直接使用
import (f "fmt")


// main函数只能声明在main包中，不能声明在其他包中
// 并且，一个main包中也必须有且仅有一个main函数

// 函数的左大括号必须和函数在同一行
func main(){
	// Println 用于格式化输出
	f.Println("你好 golang")
}





