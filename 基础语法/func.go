package main

import "fmt"

func main() {
	a, b := 1, 2
	A(a, b)
	fmt.Println("########不定长变参##########")
	fmt.Println("传入非slice时，不改变", a, b)
	fmt.Println("#############")
	fmt.Println("传入s参数为slice时会改变")
	s1 := []int{1, 2}
	fmt.Println(s1)
	B(s1)
	fmt.Println(s1)
	fmt.Println("#######匿名函数######")
	lamda()

	fmt.Println("#######闭包######")
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))

}

// 不定长参数传入
func A(a ...int) {
	a[0] = 3
	a[1] = 4
	fmt.Println(a)
}

func B(s []int) {
	for i := range s {
		s[i] = 0
	}
}

func lamda() {
	a := func() {
		fmt.Println("这是匿名函数")
	}
	a()
}

func closure(x int) func(int) int {
	fmt.Println("%p\n", &x)
	return func(y int) int {
		fmt.Println("%p\n", &x)
		return x + y
	}
}
