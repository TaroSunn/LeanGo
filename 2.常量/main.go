package main

func main() {
	// 常量使用 const 关键词定义
	const constVariables1 float64 = 3.1415

	// const 也支持多值赋值
	// iota 可以用于枚举类型
	const (
		Error   = iota // 0
		Success        // 1
	)
}
