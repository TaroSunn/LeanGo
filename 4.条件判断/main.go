package main

import "fmt"

func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	}
	return g
}

func main() {
	// 在go中 进行条件判断时是不能加括号的
	const success = true
	if success {
		fmt.Println("success")
	} else {
		fmt.Println("error")
	}

	fmt.Println(
		grade(100),
	)
}
