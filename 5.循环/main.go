package main

import "fmt"

func main() {
	// go 中的循环也不能添加 括号
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
}
