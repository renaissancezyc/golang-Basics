package main

import (
	f "fmt"
	"sort"
)


func main(){

	/*切片 slice

		1. 	切片是对数组的一个连续判断的引用，所以切片是一个引用类型（类似于
			Python中的list类型），这个判断可以是整个数组，或者是由起始和终止索引
			表示的一些项的子集，需要注意的是，终止索引表示不包括在切片内。

		2.	切片是可索引的，并且可以使用函数len()获取切片的长度,
			切片还提供了计算容量的函数cap()可以测量切片最长可以达到多少也就是容量，它等于
			切片从第一个元素开始，到相关数组末尾的元素个数，如果s是一个切片，cap(s)
			就是从s[0]到数组末尾的数组长度。切片的长度永远不会超过它的容量。

		3.	切片的初始化格式：var slice1 []type = arr1[start:end]
			这表示slice1是由arr1从start索引到end索引之间的元素构成的子集（切分数组，
			start:end被称为 slice表达式）。所以slice1[0]就等于arr1[start]。
			这可以在arr1被填充前就定义好。

		5.	如果是 var slice1 []type = arr1[:] 那么selice1就等于完整的arr1数组（
			所以这种写法其实是 arr1[0:len(arr1)]的缩写）。

		6.	切片在内存中的组织方式实际上是有3个结构体：指向相关数组的指针，切片长度以及切片容量。

		7.	注意 绝对不要用指针指向 slice。切片本身已经是一个引用类型，所以它本身就是一个指针！！，
			当然slice也不是指针，知识包含指针的结构体，可看第四条。
	
	
	*/  

	var b []int

	/*  批量添加

		append 需要先指定要添加数据的切片名称
		b = append(b, 1,3)
	*/
	for i := 0; i < 5; i++ {
		b = append(b, i)
	}
	f.Println(b)

	// 拷贝切片
	ts1 := []int{1,2,3}
	ts2 := []int{4,5,6}
	copy(ts1,ts2)  // 输出：[4 5 6] [4 5 6]
	//func copy(dst,src []T) 
	// 显而易见copy方法把ts2的切片从源地址拷贝到了目标地址,并覆盖了原来的ts1的值。
	f.Println(ts1,ts2)
	

	// 第一种遍历切片

	for i := 0; i < len(b); i++ {
		f.Println(b[i],i)
	}


	// 第二种遍历切片
	list1 := []int{11,2,3,4,5,6}
	for key,value := range(list1){
		f.Println(key,value)
	}


	// 将切片传递给函数
	test1(b)


	// 使用make()创建一个切片
		// make接收三个参数：元素的类型，长度和容量，并返回切片
		s1 := make([]int,10,20) // cap(slice2)=len(slice2) == 10
		f.Printf("s1切片的长度为：%d\ns1切片的容量为：%d\n",len(s1),cap(s1))
		// 输出：s1切片的长度为：10 ，s1切片的容量为：20

		s2 := make([]int,10)
		f.Printf("s2切片的长度为：%d\ns2切片的容量为：%d\n",len(s2),cap(s2))
		// 输出： s2切片的长度为：10 ， s2切片的容量为：10
		// 可以看到当你指定切片长度时，切片容量可以默认为切片长度


	// 使用new创建一个切片
		s3 := new([100]int)[0:50]
		f.Printf("s3切片的长度为：%d\ns3切片的容量为：%d\n",len(s3),cap(s3))
		//输出： s3切片的长度为：50 , s3切片的容量为：100 当然容量也可以不写，默认和长度一致


	
	/* make() 和 new()的区别：

		new(T) 为每个新的类型T分配一篇内存，初始化为0并且返回类型为*T的内存地址：
		这种方法返回一个指向类型T，值为0的地址的指针，它适用于值类型如数组和结构体
		相当于&T{}

		make(T) 返回一个类型为T的初始值，它只适用于3种内建的引用类型：切片、map和channel
	
	*/


	/* 搜索和排序切片
		标准库提供了sort包来实现常见的搜索和排序操作
	*/

	tt3 := []int{9,2,5,8}
	sort.Ints(tt3) // 对tt3进行了int类型的升序排序
	f.Println(tt3) // 输出：[2 5 8 9]
	

	tt4 := []string{"aaaa","bbbbb","qqqq","aaaa"}
	sort.Strings(tt4)
	f.Println(tt4) // 输出：[aaaa aaaa bbbbb qqqq]

	// 想要在数组或者切片里搜索一个元素，该数组或切片必须先排序（因为标准库的算法使用的是二分法）
	tt4_index := sort.SearchStrings(tt4,"qqqq")
	f.Println(tt4_index) // 输出3






	aa := test2("Google")
	f.Println(aa)
	
	


}

// 将切片传递给函数
func test1(arr []int){
	for key,value := range(arr){
		f.Println(key,value)
	}
}


func test2(a string) string{

	r := make([]byte,len(a))

	for i:=0;i<len(a);i++{
		r[i]=a[len(a)-i-1]
	}

	return string(r)

}
