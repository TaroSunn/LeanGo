package main

import "fmt"

func main() {
	m := map[string]string{
		"name": "sun",
	}
	m2 := make(map[int]int)
	fmt.Println(m, m2)
}
