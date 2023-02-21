/*
	在golang里如何转换一个变量的数据类型
*/

package main

import f "fmt" 

/* 
一个类型的值可以被转换成另一种类型的值，由于Go语
言不存在隐士类型转换，因此所有的转换都必须显式说明。
*/ 
func main(){

	data1 := 3.1415
	data2 := int(data1)

	f.Println("float类型转换int类型的结果:",data2)
	// float类型转换int类型的结果: 3

}