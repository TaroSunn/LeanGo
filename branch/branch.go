package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "a.txt"

	contents, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err, "123")
	} else {
		fmt.Printf("%s\n", contents)
	}

}
