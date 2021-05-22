package main

import "fmt"

func main() {
	// 定义 arr1 是包含5个int值得数组
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	// ... 表示可能有n个
	arr3 := [...]int{2, 4, 5, 6, 7}
	fmt.Println(arr1, arr2, arr3)

	// 遍历数组
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	// 使用 range 可以获取 数组下标 和数组下标对应的值
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
}
