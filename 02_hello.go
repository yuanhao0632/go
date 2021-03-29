// go语言以包为单位
// 每个文档都需要先声明包的名称
// 程序必须有一个main包
// 一个工程（文件夹）只能有一个主函数,如果在一个文件夹就用命令行执行
package main

// 倒入package,导入包之后一定要使用，否则会出现错误
import "fmt"

func main() {
	var a int
	var b, c int
	fmt.Println("hello go !!!")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)

	a = 10 //变量赋值
	fmt.Println("a = ", a)
	// 3.自动推导类型
	d := 30
	//%T打印变量所属的类别
	fmt.Printf("d type is %T\n", d)

	// 4
	fmt.Println("a = ", a)    // 一段一段处理，自动加换行
	fmt.Printf("a = %d\n", a) // 格式化输出，将a放在%d的位置，输出是a = 10\n

	e, f := 20, 30
	fmt.Printf("a = %d, e = %d, f=%d\n", a, e, f)
	// 5
	i, j := 10, 20
	i, j = j, i
	fmt.Printf("i = %d, j = %d\n", i, j)

} // go语言语句结尾是没有；分号的

/*
这是块注释
go build 02_hello.go  对go进行编译，编译为可执行程序02_hello后可以直接在当前目录用./02_hello运行这个可执行文件

go run XXX.go  不生成可执行程序就运行

1.变量声明 var 名字 类型，定义对变量一定要使用，否则会报错
只是声明，没有初始化对变量，默认为0
同一个{}里声明对变量名是唯一的，
可以同时声明多个变量
如var a int

2.变量初始化，实际上就是声明变量的同时对变量进行赋值
var b int = 10 // 初始化，声明和赋值一起
b = 20 // 赋值，需要先声明再赋值

3.自动推导类型，必须初始化，通过初始化的值确定类型(常用写法)
自动推导对于同一个变量名只能使用一次（b := 30），就是初始化那次，但是赋值可以使用很多次(b = 30)

4.Println和Printf的区别

5.多重赋值和匿名变量
*/
