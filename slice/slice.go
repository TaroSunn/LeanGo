package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6]
	fmt.Println(s)

	updateSlice(s)
	fmt.Println(s, arr)

	s2 := append(s, 10)
	s3 := append(s2, 11)
	fmt.Println(s2, s3)
}
