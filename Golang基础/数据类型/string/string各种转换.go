/*
	1.HasPrefix 	判断字符串是否以指定字符开头
	2.HasSuffix 	判断字符串是否以指定字符结尾
	3.Contains  	判断字符串是否包含指定字符
	4.Index         获取指定值在字符串中第一次出现位置的下标
	5.LastIndex 	获取指定值在字符串中最后出现位置的下标
	6.Replace		替换指定old字符为new字符，并且传递替换次数-1就是替换所有
	7.Count			统计字符出现次数
	8.Repeat 		重复字符串，并返回一个新的字符串
	9.ToLower		将字符串中的Unicode字符全部转换为相应的小写字符
	10.ToUpper		将字符串中的Unicode字符全部转换为相应的大写字符
	11.TrimSpace 	剪切字符串开头和结尾的空白符号，剪切指定字符使用Trim，如果想要指定只剪切开头或结尾则使用，TrimLeft 或者 TrimRight。
	12.Fields		利用空白作为分隔符将字符串分割为若干块，并返回slice。指定字符分割使用Split也返回slice
	13.Join			将元素类型为string的slice使用分割符号里拼接成一个字符串
	14.NewReader  	从字符串中读取内容，用于生成一个Reader并获取字符串中的内容，返回Reader的指针。

	字符串相关数据类型转换都是以通过strconv包实现的。
	15.strconv.Atoi  strconv.Itoa 字符串与数字之间的相互转换
周奕呈-
*/

package main

import (
	f "fmt"
	"strconv"
	"strings"
)

// 作为一种基本数据类型，每种语言都有一些对于字符串的预定义处理函数
func main() {
	data1 := "Hello golang."

	// 1. HasPrefix 判断字符串是否以指定字符开头。  bool
	result1 := strings.HasPrefix(data1, "Hello")
	f.Println("result1=data1是否以Hello这个单词开头:", result1)

	// 2. HasSuffix 判断字符串是否以指定字符结尾。  bool
	result2 := strings.HasSuffix(data1, "golang")
	f.Println("result2=data1是否以golang这个单词终止:", result2)

	// 3. Contains 判断字符串是否包含指定字符。  bool
	result3 := strings.Contains(data1, "ello")
	f.Println("result3=data1里是否包含ell这三个字符:", result3)

	// 4. Index 在字符串中查找指定值（第一次出现的指定值）的下标 。 int 返回指定值下标值或-1，-1代表不存在。
	result4 := strings.Index(data1, "golang")
	f.Println("result4=golang在变量data1里的下标是:", result4)

	// 5. LastInde 获取指定字符最后一次出现的下标。int 返回指定值下表或-1。-1代表不存在。
	result5 := strings.LastIndex(data1, "l")
	f.Println("result5=data1里最后一次l字符的下标位置是:", result5)

	// 6. Replace 替换指定old字符为new字符，并且传递替换次数-1就是替换所有
	// func Replace(s, old, new string, n int) string
	result6 := strings.Replace(data1, "golang", "go", 1)
	f.Println("result6=将data1里的golang替换为go,并且只替换一次的结果为:", result6)

	// 7. Count 统计字符出现次数
	result7 := strings.Count(data1, "l")
	f.Println("result7=在data1中l出现的次数为:", result7)

	// 8. Repeat 重复字符串，并返回一个新的字符串
	result8 := strings.Repeat(data1, 3)
	f.Println("result8=将data1重复6次:", result8)

	// 9. ToLower 将字符串中的Unicode字符全部转换为相应的小写字符
	result9 := strings.ToLower(data1)
	f.Println("result9=data1中的字符全部转换为小写的结果为:", result9)

	// 10. ToUpper 将字符串中的Unicode字符全部转换为相应的大写字符
	result10 := strings.ToUpper(data1)
	f.Println("result10=data1中的字符全部转换为大写的结果为:", result10)

	// 11. TrimSpace 剪切字符串开头和结尾的指定字符，剪切指定字符使用Trim，如果想要指定只剪切开头或结尾则使用，TrimLeft 或者 TrimRight。
	result11 := strings.TrimSpace(data1)
	f.Println("result11=剪切data1中的指定字符:", result11)

	// 12. Fields 利用空白作为分隔符将字符串分割为若干块，并返回slice。指定字符分割使用Split也返回slice
	result12 := strings.Fields(data1)
	f.Println("以空格将data1分割为:", result12)
	// 因为是切片所以使用range进行处理
	// for _,val := range result12{
	// 	f.Println("以空格将data1分割为:",val)
	// }

	// 13. Join	将元素类型为string的slice使用分割符号里拼接成一个字符串
	result13 := strings.Join(result12, ",")
	f.Println("将分割过后的字符串以,拼接后的值为:", result13)

	// 14. NewReader() 从字符串中读取内容，用于生成一个Reader并获取字符串中的内容，返回Reader的指针。
	result14 := strings.NewReader(data1)
	f.Println("读取data1字符串的长度", result14.Len())

	// 15. strconv.Atoi  strconv.Itoa 字符串与数字之间的相互转换

	// 字符转数字
	result15, _ := strconv.Atoi("10") // strconv.Atoi 返回两个数据(i int,err error)
	f.Println("字符串result15转换成数字后的值为:", result15)

	// 数字转字符
	result_15_int := result15 + 5
	result_15_string := strconv.Itoa(result_15_int)
	f.Println("result_15转换成+5后的字符串为:", result_15_string)

	/* 打印结果
	result1				data1是否以Hello这个单词开头: true
	result2				data1是否以golang这个单词终止: false
	result3				data1里是否包含ell这三个字符: true
	result4				golang在变量data1里的下标是: 6
	result5				data1里最后一次l字符的下标位置是: 8
	result6				将data1里的golang替换为go,并且只替换一次的结果为: Hello,go.
	result7				在data1中l出现的次数为: 3
	result8				将data1重复6次: Hello,golang.Hello,golang.Hello,golang.
	result9				data1中的字符全部转换为小写的结果为: hello,golang.
	result10			data1中的字符全部转换为大写的结果为: HELLO,GOLANG.
	result11			剪切data1中的指定字符: Hello golang.
	result12			以空格将data1分割为: [Hello golang.]
	result13			将分割过后的字符串以,拼接后的值为: Hello,golang.
	result14			读取data1字符串的长度 13
	result15			字符串result15转换成数字后的值为: 10
	result_15_string	result_15转换成+5后的字符串为: 15
	*/

}
