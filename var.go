package main

import (
	"fmt"
	"math"
)

//别名
type (
	文本 string
)

/*变量的声明与赋值：
变量的声明格式：var <变量名称> <变量类型>
变量的赋值： var <变量名称> = <表达式>
声明的同时赋值： var <变量名称> <变量类型> = <表达式>
*/
//全局变量
var (
	a         [3]int32
	b1, _, b3 int = 1, 2, 3
	text      文本
)

// 类型不赋值时输出类型的#0值#
func main() {
	// ":="代替了var声明并赋值，只能在方法体中使用
	c := 65.7
	//类型转换必须显示声明，只能转换兼容类型
	d := int(c)
	text = "中文类型别名so cool!但最好别这样弄"
	fmt.Println(math.MinInt32)
	fmt.Println(a, b1, b3, c, d, text)

}
