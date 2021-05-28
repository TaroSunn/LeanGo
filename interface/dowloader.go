package main

import (
	"fmt"
	"interface/infra"
)

func getRetriever() infra.Retriever {
	return infra.Retriever{}
}

func main() {
	retriever := getRetriever()

	fmt.Println(retriever.Get("https://baidu.com"))
}
