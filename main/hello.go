package main

import "fmt"
import "unsafe"

// 单行注释
/*
 Author by 菜鸟教程
 我是多行注释
*/

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

func main() {
	fmt.Println("Hello, World!")
	// 字符串连接
	fmt.Println("Google" + "Runoob")

	// %d 表示整型数字，%s 表示字符串
	var stockCode = 123
	var endDate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var targetUrl = fmt.Sprintf(url, stockCode, endDate)
	fmt.Println(targetUrl)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	var a string = "Runoob"
	fmt.Println(a)

	var bb, c = 1, 2
	fmt.Println(bb, c)

	/* 以下是go语言变量学习 */

	/* 以下是go语言常量学习 */
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const aaa, bbb, ccc = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(aaa, bbb, ccc)

	// --------------------------------------
	/* 以下是if语句学习 */

}
