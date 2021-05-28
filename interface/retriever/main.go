package main

import (
	"fmt"
	"interface/retriever/mock"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is web"}
	fmt.Println(download(r))
}
