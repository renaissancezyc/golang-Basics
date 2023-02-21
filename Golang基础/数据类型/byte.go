package main

import (
	f "fmt"

)

// 字符
func main(){
	// 字符是组成字符串的最小元素，golang中使用byte定义字符串，
	// byte表示utf-8字符串单个字节的值

	// byte 使用单引号 '' 代表字节码
	// 双引号 "" 代表字符串

	var x byte = 65
	var y uint8 = 65

	f.Printf("x = %c\n",x) // x = A
	f.Printf("y = %c\n",y) // y=A
	// x = A
	// y = A

}