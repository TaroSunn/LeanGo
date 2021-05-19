package main

import "fmt"

var (
	variables3 int
)

func main() {
	// 变量声明 使用关键词 var
	// var关键字 + 变量名 + 类型
	// 类型 也有可无
	// 如果变量赋值那么会有默认值,例如 int类型 默认为 0
	var variables int
	// 除了上面一种方式声明变量，另一种声明变量的另一种方式是
	// 变量名 := 值
	// 这种方式只能在函数内部使用
	variables2 := "Hello"
	fmt.Println(variables, variables2)

	// 在go中可以 一次定义赋值多个变量
	var a, b, c = 1, "123", 3
	fmt.Println(a, b, c)
}
