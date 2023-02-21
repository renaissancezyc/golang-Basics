package main

import (
	f "fmt"
)

/*
	1. 指针就是地址。 指针变量就是存储地址的变量。一个程序在内存中存储它的值，每个内存块(或字)有一个地址，通常用16进制数表示，如：0x6b0987
	或者 0xf98001d7f0。

	2. 一个指针变量可以指向任何一个值的内存地址，它指向那个值的内存地址，在32位机器上占用4个字节，
	在64位机器上占用8个字节，并且与它所指向的值的大小无关。也可以声明指针指向任何类型的值来
	表明它的原始性或结构性，可以在指针类型前面加上*号来获取指针所指向的内容，这里的*号是一、
	个类型更改器，使用一个指针引用一个值被称为间接引用。

	3. 一个指针被定义后没有分配任何变量时，值为nil。

	4. 一个指针变量通常缩写为ptr。

	5. *p： 解引用、间接引用

	6. & 关键字可以从一个变量中取到其内存地址。

	7. * 关键字如果在赋值操作值的左边，指该指针指向的变量；* 关键字如果在赋值操作符的右边，指从一个指针变量中取得变量值，又称指针的解引用。
	
	8. 此外，指针也可以拥有指针，也就是说，指针也可以指向别的指针所在的内存

*/


func main(){

	
	/*  例子1：
			这个内存地址可以存储在一个叫指针的特殊数据类型中,
			定义一个指向int的指针，此时使用zz=&data1，就代表zz指向data1。
			zz存储了data1的内存地址，zz指向了data1的位置。
	*/
	data1 := 66
	f.Printf("这是data1的内存地址：%p\n",&data1) //这是data1的内存地址： 0xc000014088
	var zz *int 
	zz = &data1
	f.Printf("zz的值是什么：%p，zz指向的变量为：%d",zz,*zz)  // zz的值是什么：0xc0000aa058，zz指向的变量为：66

	// 例子2：

	aa := "Hi,Renaissance"
	bb := &aa
	*bb ="Hi,zhouyicheng"

	f.Println("aa字符串改变后的值为：",aa) // aa字符串改变后的值为： Hi,zhouyicheng
	f.Println("bb字符串改变后的值为：",*bb) //  指针bb字符串指向的值为： Hi,zhouyicheng

	// 注意你不能得到一个文字或者常量的地址
	const cc = 888

	ptr := &cc 
	ptr3 := &10
	
	f.Println(ptr) // invalid operation: cannot take address of cc
	f.Println(ptr3) // invalid operation: cannot take address of 10



}