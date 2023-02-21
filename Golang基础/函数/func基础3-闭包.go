package main

import f "fmt"


func main(){

	// 1. 将函数作为参数传递 调用回调函数
		callback(test1)
		// 输出：我是test1，a=1,b=3

	// 2. 匿名函数
		test2()
		// 输出：我是test2里的匿名函数： 8

	// 3. 闭包操作，将函数作为返回值
		t3 := test3()
		f.Println("我是test3的返回值：",t3(1))
		// 输出 我是test3的返回值： 2     

		t4 := test4(5)
		f.Println("我是test4的返回值：",t4(5))
		// 输出 ：我是test4的返回值： 10   



}


// 1. 将函数作为参数传递
func test1(a,b int){
	f.Printf("我是test1，a=%d,b=%d\n",a,b)
}

func callback(f func(int,int)) {
	f(1,3)
}

/* 2. 匿名函数
	当我们不想给函数起名字的时候，可以使用匿名函数。
	例：func(x,y int) int {return x+y}。
	但是这样的一个函数不能够独立存在，需要赋值给某个变量然后通过调用变量名
	来调用这个函数，当然也可以直接调用,
	例：func(x,y int) int {return x+y} (3,4)
*/

func test2(){

	// 将匿名函数赋值给变量tt
	tt := func(i int){f.Println("我是test2里的匿名函数：",i)}
	// 给变量tt传参
	tt(8)

}


/* 3. 闭包操作，将函数作为返回值
	假设有两个函数，
	func test3() (func (b int) int )  
	函数test3不接收参数，返回值是一个函数，test3返回的函数接收一个int类型的参数，test3返回的函数又返回一个int数据。
	
	func test4(a int) (func (b int) int)
	函数test4接收一个int类型的参数，返回值是一个函数，test4返回的函数接收一个int类型的参数，test4返回的匿名函数又返回一个int数据。

*/

/* 函数名为test3，test3的返回值为一个函数 ，要返回的函数接收int类型参数，要返回函数的参数为int类型的数据。
有点小饶多看几遍理解一下，相当于你调用test3，test3又返回给你另一个函数，之后你给返回回来的函数再传参，最后
函数返回值。
*/

func test3() func(a int) int{
	return func (a int) int {
		return a+1
	}
}

func test4(a int) func(b int) int{
	return func (b int) int {
		return a+b
	}
}