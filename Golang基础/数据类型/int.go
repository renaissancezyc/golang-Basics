package main

import (
	f "fmt"
	"math"
)

// 有符号整形
func main(){
	// 整形顾名思义就是存储数据类型是正数。
	// golang中分为有符号和无符号，简单理解就是存储范围上的差异，
	// 有符号整形分为 int 86 32 64  以及默认的int
	// 除非我们需要使用特定大小的整数，否则通常都使用int来表示整数
	var num8 int8 = 127
	var num16 int16 = 32767
	var  num32  int32 = math.MaxInt32
	var num64 int64 = math.MaxInt64
	var num int = math.MaxInt

	// golang 支持累加操作  a:=1  a++
	// python 就不支持累加操作 必须这样写  a+=1

	f.Printf("num8是 %d\n",num8)
	f.Printf("num16是 %d\n",num16)
	f.Printf("num32是 %d\n",num32)
	f.Printf("num64是 %d\n",num64)
	f.Printf("num是 %d\n",num)

}