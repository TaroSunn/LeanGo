package main

import "fmt"

func main() {
	m := map[string]string{
		"name": "abc",
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
