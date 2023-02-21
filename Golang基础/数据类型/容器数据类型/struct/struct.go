package main

import (
	f "fmt"
)


// 结构体(struct)  可以把结构体理解为python里的class
// 结构体不存在向对象语言里
// 结构体以类型的形式存在


// 声明结构体 人类结构体
type Human struct {

	// 定义结构体属性 可以理解为python类里的self属性
	// 定义属性必须属性的数据类型
	username string  // 名字
	age int // 年龄

}

// 学生结构体
type Student struct {

	// 包含人类结构体  当在golang里想要写继承的时候
	// 那就写成包含的写法 比如学生继承人类 那就是学生包含人类
	human Human
	class string

}


func main(){

	// // 如何使用结构体
	// // 创建结构体  在python里可以理解为实例化类
	// human := Human{
	// 	username:"小明",
	// 	age:10,  // 如果实例化结构体的时候，并没有给age传值 那age的数据根据数据类型赋空值 
	// }
	
	// // 打印结构体 
	// f.Println(human)
	// // 也可以访问结构体中的某个属性
	// human.age = 30
	// f.Println(human.age)
	// // 结构体解决了map容器数据类型的问题  map容器数据类型所有的内容类型必须一样
	// // 而结构体允许不同的数据类型，存在结构体中
	// // 结构体也可以动态的修改结构体属性


// __________________________________________________________
	// 指针调用结构体
	// 
	pointer_test := &Human{username:"我是结构体指针测试数据",age:666}
	f.Println("&Human ==",(*pointer_test).username)
	// 打印结果: &Human == 我是结构体指针测试数据

// __________________________________________________________


	// // 结构体的包含

	// student := Student {

	// 	class: "人工智能",

	// }
	// // 我们可以理解为 Student结构体中包含了Human结构体 
	// // 那我们就可以使用Human结构体里的属性 给属性定义值

	// // 可以这样定义student里的human属性
	// student.human = Human {

	// 	username:"小明",
	// 	age:20,

	// }

	// // 也可以直接把上面定义的Human赋值给studnt的human
	// student.human = human  //这里的等于human就是等于第28行的human


	// f.Println("9999",student)




	
	
}