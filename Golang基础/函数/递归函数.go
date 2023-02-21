package main

import (
	f "fmt"

	"golang.org/x/exp/slices"
)

/*
	当一个函数在其函数体内调用自身，则称之为递归。最经典的例子就是计算斐波那契数列,
	既前两个数相加等于第三个数 1 1 2 3 5 8 13 21 34 55 89 144
*/

func main(){

	//  1. 递归实现斐波那契数列
	for i:=0;i<=10;i++{
		f.Println(test1(i))
	}
	// 输出：1 1 2 3 5 8 13 21 34 55 89

}


// 1. 递归实现斐波那契数列
func test1(oo int) (res int ) {
	if oo<=1{
		res = 1
	}else{
		res = test1(oo-1) + test1(oo-2)
	}

	return 
}

/* 2. 栈溢出问题
		一般出现大量的递归调用导致程序栈内存分配耗尽，这个问题可以通过一个名为
		懒惰求值的技术解决，在Go语言中，我们可以使用管道（chanel）和（goroutine）
		来实现。
*/ 