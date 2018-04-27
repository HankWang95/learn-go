package main

import (
	"fmt"
)

func _array() {
	fmt.Println("数组学习")
	//必须声明数组长度，不同长度的数组不是一个类型，不能够比较
	var b [1]int
	var c [1]int
	b[0] = 1
	c[0] = 1
	fmt.Println(b == c)
	//数组作为参数传值时不是传的地址，而是整个数组数列，所以可以用数组指针的概念
	a1 := [2]int{1, 2}
	a2 := [2]int{1, 2}
	//声明数组指针的两种形式，
	var p1 *[2]int
	p2 := new([2]int)
	p1 = &a1
	p2 = &a2
	fmt.Println(p1, p2, p1 == p2)
	//数组索引
	d := [10]int{9: 1}
	fmt.Println(d)

}

func blublu() {
	a := [...]int{1, 2, 3, 5, 6, 9, 13, 7}
	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println("冒泡排序")
	fmt.Println(a)
}

func main() {
	blublu()
	_array()
}
