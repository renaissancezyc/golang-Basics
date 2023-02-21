package main


import (
	f "fmt"
	"reflect"
)

func main(){

	// 1.传递变长参数
	test1("aaaaa","bbbbb","ccccccc","dddddd")
	/* 输出:
		a的值为： aaaaa
		b的值为： [bbbbb ccccccc dddddd]
	*/

	// 2.结构体传参
	test2(1,2,Options{3,"hhhh",88.88})
	/* 输出:
		a的值为： aaaaa
		b的值为： [bbbbb ccccccc dddddd]
	*/

	// 3.使用空接口传参
	test3(11111)


}


/* 1.传递变长参数
		如果函数的最后一个参数是采用...type的形式，那么这个函数就可以处理一个变长
	参数，这个长度可以为0，这样的函数被称为变长函数。如果想要传递的参数为slice，则可以通过slice...
	来传递参数
*/
func test1(o1 string,o2 ...string){
	f.Println("a的值为：",o1)
	f.Println("b的值为：",o2)
}



/* 2.使用结构传参
		如果变长参数的类型并不是都相同的该怎么办,这时候就使用结构体传参
*/

type Options struct {
	par1 int
	par2 string
	par3 float64
}

func test2(a int ,b int ,c Options){ // 也可以提前初始化好结构体 func test2(a,b,Options{par1:1,par2:"qqqq",par3:52.55})
	f.Println("a的值为:",a)
	f.Println("b的值为:",b)
	f.Println("options的值为:",c)
}


/* 3.使用空接口
		如果变长参数的类型没有被指定，则可以使用默认的空接口interface{}，这样
		就可以接收任何类型的变长参数，一般使用for——range或者switch来判断
*/

func test3(data ...interface{}){
	for _,value := range data{
		switch uu := reflect.TypeOf(value){
			case int: f.Println("整型")
			case float64: f.Println("浮点型")
			case string: f.Println("字符串")
			default : f.Println("检测不出来")
			// .....
		}
	}
}
