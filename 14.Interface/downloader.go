package main

import (
	"fmt"
)

func main() {
	retriverer := infra.Retriever{}

	fmt.Println(retriverer.GET("https://baidu.com"))
}
