package main

import (
	"fmt"
)

//指针
func _zhizhen() {
	a := 1
	//指针
	var p0 *int = &a
	fmt.Print(p0)

	var b [3]int
	//数组
	b = [3]int{1, 2, 3}
	//数组指针
	var p1 *[3]int = &b
	//二级指针
	var p2 **[3]int = &p1
	fmt.Print(b, p1, p2)

}

//if
func _if() {
	a := 10
	// if 中可以定义局部变量，if语句块结束后释放
	if a, b := 1, 2; a > b {
		print("a>b")
	} else {
		println("a<b")
	}
	// 这里的a还是第27行的
	print(a)

}

//循环——只有for
func _for() {
	a := 1
	// 不加任何语句的for就是无限循环，没有break就一直转
	for {
		a++
		if a > 3 {
			println("跳出无限循环")
			break
		}
	}
	//for后加条件，如果满足就转
	for a <= 4 {
		a++
		println("a =", a)
	}
	//for 中初始化值和步进表达式
	for i := 0; i < 3; i++ {
		a++
		println(a)
	}

}

//switch:pass

//goto break continue
func _break_goto_continue() {
	println("###########break##########")
LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 2 {
				// 直接跳出 LABLE 所在循环
				println("跳出LABLE循环")
				break LABEL1
			}
		}
	}
	println("跳出到这里")
	//break到这里继续执行
	println("##########continue###########")
LABEL2:
	// continue 到这里继续执行
	for i := 0; i < 10; i++ {
		for {
			if i > 3 {
				//
				goto LABEL3
			}
			println("跳至LABLE2循环")
			continue LABEL2

		}
	}
LABEL3:
	print("goto 到这里了：goto的lable可以不是循环，break和continue必须是循环")

}

func main() {
	_for()
	_if()
	_break_goto_continue()
	_zhizhen()
LABLE4:
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			i++
			continue LABLE4
		}
	}
}
