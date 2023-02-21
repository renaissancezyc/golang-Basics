package main

import (
	f "fmt"
)




func main(){

	map1 := map[string]int{
		"你好":123456,
		"hh":666666,
		"博客":129087,
		"Golang":666666,
	}

	/* 1. 判断key是否存在map中：
		value,ok := map1["hh"],如果hh这个键存在的话ok返回true，否则返回false
	
	*/
	_,ok := map1["hh"]
	f.Println(ok)  // 输出：true

	if _,ok := map1["yyyyy"]; ! ok{
		f.Println(ok) // 输出：false
		f.Println("map1中不存在名为yyyyy的key")
	}


	// 2. 从map1中删除一个key，如果这个key不存在，该操作也捕获产生错误
	delete(map1,"hh")  // 删除map1字典里key为hh的数据
	f.Println(map1)  // 输出：map[Golang:666666 你好:123456 博客:129087]

	delete(map1,"oo") // 删除一个不存在于字典的key
	f.Println(map1) // 输出：map[Golang:666666 你好:123456 博客:129087] 可以看到是正常输出，没有报错



	// 3. 循环获取map中的数据
	for key,value := range map1{
		f.Printf("map1中的key:%s,value:%d\n",key,value)
	}
	/* 输出:
		map1中的key:你好,value:123456
		map1中的key:博客,value:129087
		map1中的key:Golang,value:666666
	
	*/


	/*
		4. map的排序
			在map中默认是无序的，不管按照key还是按照value默认都不排序。
			如果你想为map排序，需要将key（或者value）拷贝到一个切片里（
			使用sort包进行排序）
	*/




}