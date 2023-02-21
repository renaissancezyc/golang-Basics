package main

import (
	f "fmt"
	"math"
)

// 浮点
func main(){
	// 浮点数据类型表示存储的数据是实数，比如说3.1415926，主要分为32和64位
	// 区别在于长度。和整数一样，除非需要使用特定大小的浮点，否则通常都使用float来表示浮点 

	var num1 float32 = math.MaxFloat32
	var num2 float64 = math.MaxFloat64

	f.Printf("num1的类型是%T,num1是%g\n",num1,num1)
	f.Printf("num2的类型是%T,num2是%g\n",num2,num2)

	// num1的类型是float32,num1是3.4028235e+38
	// num2的类型是float64,num2是1.7976931348623157e+308

}