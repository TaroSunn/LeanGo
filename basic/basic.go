package main

import "fmt"

var aa = 3

func variable() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s) // 0 ''
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, s, b)
}

func variableTypeDecution() {
	var a, b, c = 3, 4, true
	var s = "abc"
	fmt.Println(a, b, c, s)
}

func varialbeShorter() {
	a, b, c := 3, 4, true
	b = 2
	fmt.Println(a, b, c)
}

func main() {
	fmt.Println("hello world")
	variable()
	variableInitialValue()
	variableTypeDecution()
	varialbeShorter()
}

// 使用 var 关键字 定义变量
// 可以定义多个变量
