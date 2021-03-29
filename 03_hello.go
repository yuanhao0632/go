// 匿名变量
package main 

import "fmt"

// go函数可以返回多个值
func test() (a, b, c int) {
	return 1, 2, 3
}

func main() {
	i, j := 10, 20
	fmt.Printf("i = %d, j = %d\n", i, j)

	var tem int 
	tem, _ = i, j // 这里_是匿名变量，会丢弃掉j
	fmt.Printf("tem is %d\n", tem)
	// 匿名变量配合函数返回值使用才有优势
	var a, b, c int
	a, b, c = test()
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	a, _, _ = test()
	fmt.Printf("a = %d\n", a)

	// 常量的使用
	// 常量：程序运行期间不可以改变的量，常量声明需要const关键字
	const d int = 10
	// d = 20 // 常量不允许修改
	fmt.Printf("d = %d\n", d)
	// 常量的自动推导类型
	const e = 10 // 没有使用：=
	fmt.Printf("e = %d\n", e)
	fmt.Printf("e type is %T\n", e)
}

