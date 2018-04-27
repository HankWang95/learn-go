package main

import (
	"fmt"
)

func main() {
	// 声明一个map
	// var m map[int]string
	// 创建一个m的实例
	// m := make(map[int]string)
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m))
	fmt.Println(m)
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	// map类似python 的字典类型，hash表，每次输出都有可能不一样
	fmt.Println(s)
	home_work()
}

func home_work() {
	// 把map[int]string 的键和值进行交换
	src_m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	res_m := make(map[string]int)
	for k, v := range src_m {
		res_m[v] = k
	}
	fmt.Println(res_m)
}
