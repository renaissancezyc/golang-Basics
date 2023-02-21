package main

import (
	f "fmt"
)

// for 循环
// Golang中只有一种循环结构: for循环
func main (){

	// for 初始化语句；条件语句；修饰语句{}


	// 普通循环
	sum :=0
	for i:=0; i<10; i++{
		sum += i
	}


	// 初始化语句和后置语句是可选的
	sum1:=1
	for ;sum1 < 1000;{
		// f.Println(sum1)
		sum1+=sum1
	}

	// for循环也相当于Golang里的while(无限循环)，在上一个for循环下把分号去除即可
	sum2:=1
	for sum2 < 1000{
		// f.Println(sum2)
		sum2+=sum2
	}

	f.Println(sum1)
	/*
		for-range结构:
			这是Go特有的一种迭代结构，它可以迭代任何一个切片，数组
	
	*/


	/* continue关键字 忽略剩余的循环体直接进入下一次循环的过程，但不是无条件执行下一次循环
	执行之前仍旧需要满足循环的判断条件 
	*/
	for i:=0;i<=5;i++{
		if i==2{
			continue
		}

		f.Println(i)
		
	}
	// 输出 0 1 3 4 5 


	//  break关键字，对条件进行检查，判断是否需要停止这层循环

	for i:=0;i<=5;i++{
		if i==2{
			break
		}

		f.Println(i)
		
	}
	// 输出 0 1

	/*	goto语句可以无条件的转移到过程中指定的行,
		通常与条件语句配合使用，可以实现条件转移，构成循环，跳出循环等。
		在结构化程序设计中一般不主张使用goto语句，一面造成程序流程混乱。
		goto对应标签既可以定义在for循环前面，也可以定义在for循环后面，
		当跳转到标签地方时，继续执行下面的代码。
	*/

	TEST:
	for i:=0;i<=10;i++{
		if i==3{
			goto TEST
		}
		f.Println(i)
	}

	// 输出：  一直循环 0 1 2 0 1 2 



	/*
		goto语句本身就是做跳转用的，而break和continue是配合for使用的。所以goto的使用不限于for，
		通常与条件语句配合使用，在for循环里break和continue也可以配合标签使用
	*/

	


}