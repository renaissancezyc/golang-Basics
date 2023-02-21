package main

import (
	f "fmt"
)

/* 1.定义函数
golang内部不允许设置默认值  n=0
可以提前在外面使用var定义
只有当某个函数需要被外部包调用的时候在使用大写字母开头
*/
func sum_n(n int) int {

	if n == 0 {
		return 0
	}

	return n + sum_n(n-1)

}

/* 2.当连续两个或多个函数的已命名形参类型相同时,除最后一个类型以外,其他都可以省略
函数也可以返回任意数量的返回值*/
func test(x,y string) (string,string){
	
	return x,y

}



/* 3.Go语言的返回值可以被命名,它们会被视作定义在函数顶部的变量
没有参数的return语句返回已经命名的返回值，也就是直接返回 
*/
func test2 (num int)(x,y int){

	x = num-1
	y = num-2
	return 


}

/* 4. 空白符
	空白符用来匹配一些不需要的值，然后丢弃掉。
*/
func test3 ()(int,int,int){
	return 1,2,3
}

/* 5. 函数修改外部变量(传递指针给函数)
*/

func test4 (ptr *int){
	 *ptr = 100
}



/* main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数
 如果有 init () 函数则会先执行该函数）。如果你的 main 包的源代码没有包含 main 函数，
 则会引发构建错误 undefined: main.main。main 函数既没有参数，也没有返回类型（与 C 家族
 中的其它语言恰好相反）。如果你不小心为 main 函数添加了参数或者返回类型，将会引发构建错误：
 */


// 在程序开始执行并完成初始化后，第一个调用(程序的入口点)的函数就是main.main()
func main() {

	// res := Sum_n(5)
	// f.Println(res)
	
	// 函数返回多个返回值
	a,b := test("hello","golang")
	f.Println(a,b) // result: hello golang

	// 函数返回值命名
	f.Println(test2(6)) //result: 5 4  

	// 空白符的使用
	var t1 int
	var t3 int
	t1,_,t3= test3()
	f.Println(t1,t3) //result: 1 3


	// 直接修改变量(传递指针给函数)
	var t4 = 99
	data4 := &t4
	test4(data4) // 调用函数修改t4的参数
	f.Println(t4)

	/* 匿名函数
			在定义时调用匿名函数
			匿名函数可以在声明后调用，例如：
			func(data int) {
				fmt.Println("hello", data)
			}(100)
			注意第3行}后的(100)，表示对匿名函数进行调用，传递参数为 100。
	
	*/



}




