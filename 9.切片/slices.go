package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	// 切片 从包括2 - 3 不包括3
	s := arr[2:3]
	fmt.Println(s)
}
