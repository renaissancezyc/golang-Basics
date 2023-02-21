package main

import (
	f "fmt"
)


func main(){
	// 容器数据类型
	//  字典也叫集合 (map)
	// "key:value  key只能是可哈希的值(不可以扩容缩容的值)"


	s_dict := map[string]int{
		"小明":10,
		"小红":20,
	}

	f.Println(s_dict["小红"])//通过key取值
	// 如果要获取的key不存在的话 会返回你当前定义的value的数据类型的0值
	// 而python的话获取不存在的key会报错，或者使用python字典.get('小红',None)设定默认值


	d_dict := map[int]int{}  //允许空值

	// 赋值操作 循环 或者根据key来赋值
	for i := 0; i < 5; i++ {
		d_dict[i]=i
	}

	// 遍历操作  如果你用不到key或者value可以使用_占位符代替,_只做占位符做用并不是声明。
	for key,value := range  d_dict{


		f.Println(key,value)
	} 

	f.Println(d_dict)



	// 用切片作为map的value
	slice1 := []int{1,2,3,4,5,6,7,8,9,10}

	map99 := map[string] []int{
		"我是map3的key": slice1,
	}

	f.Println(map99)




}