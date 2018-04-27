package main

import "fmt"

func main() {
	// defer 在程序出现严重错误的时候也会运行，在函数体执行结束
	// 按照调用顺序的相反顺序逐个执行
	// 常用于资源清理、文件关闭、解锁以及记录时间等操作
	// 通过匿名函数配合可以在return之后修改函数计算结果。
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}

	fmt.Println("#########panic###########")
	A()
	B()
	A()

}

func A() {
	fmt.Println("Func A")
}

func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in b")
		}
	}()
	panic("Panic in B")

}
