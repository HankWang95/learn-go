package main

import "fmt"

// [29:55] copy()与slice

func main() {
	//创建slice，长度为3 容量为10 ，
	s := make([]int, 3, 5)
	fmt.Printf("slice长度为 %d，容量为 %d，地址为 %p\n", len(s), cap(s), &s)
	s = append(s, 1, 2, 2, 3, 2, 3, 3, 3, 3, 3, 3)
	fmt.Printf("s变成了%v，容量为%d，地址变成了%p\n", s, cap(s), &s)

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("##############reslice##############")
	// 创建一个slice，直接指向底层数组的地址
	fmt.Println("a的数组为-----", a)
	s1 := a[1:3]
	// 创建一个resilce，指向的其实是同一个数组。相对s1的索引号
	s2 := s1[1:3]
	fmt.Println("s1 := a[1:3]-----", s1, "\ns2 := s1[1:3]-----", s2)

	// copy
}
