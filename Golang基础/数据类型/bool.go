package main

import (
	f "fmt"
)

// 布尔 bool
func showBool(){
	// bool通常情况下表示真或者假通常出现在条件语句中，在Python中布尔数据类型可以参与数值运算，也可以和其他数据类型进行转换
	//  但是在golang中镇值用true表示，并且不与1相等，同样假值用false表示，并且不与0相等，使用上相对严格，没有Python那么挥洒自如
	a := true
	b := false
	f.Println("a=",a)
	f.Println("b=",b)
	f.Println("true && false = ",a && b)
	f.Println("true || false = ",a || b)

}


func main(){
	showBool()
	
}