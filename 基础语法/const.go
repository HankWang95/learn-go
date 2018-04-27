package main

import (
	"fmt"
)

//常量，不想被外部调用时加个下划线
const (
	A = 'A'
	//未定义的常量将使用上行表达式
	B
	//iota 枚举，常量组每增加一个随之增加1
	_C = iota
	D
)

//位运算符
func yun_suan_fu() {
	//左移
	a := 1 << 10
	/*
		6:0110
		11:1011
		---------
		与& 0010 = 2
		或| 1111 = 15
		异或^ 1101 =13
		&^ 0100 = 4
	*/
	b, c, d, f := 6&11, 6|11, 6^11, 6&^11
	// ^
	fmt.Println(a, b, c, d, f)
}

func main() {
	fmt.Println(A, B, _C, D)
	fmt.Println("运算符：")
	yun_suan_fu()
}
