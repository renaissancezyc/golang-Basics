package main

import f "fmt"

/* Golang利用闭包实现斐波那契数列，斐波那契主要就是前两个数相加等于第三个数，我们要的也就是第三个数怎么求呢？
	很简单前两个数相加就等于第三个数了。
*/

func main(){
	tt := test1()
	for i:=0;i<=10;i++{
		f.Println(tt())
	}


}

func test1() func() int {
	var x,y int =1,1  // 第一位等于-1  第二位等于1

	return func() int  {
		x,y = y,x+y  // 第一位等于第二位也就是1  第二位等于第三位也就是-1+1=0  以此类推一直返回我们需要的第三个数

		return y
		
	}
}