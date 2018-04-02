package main

import (
	"fmt"
	"unicode"
)

func main() {
	var src, dst string
	src = `SELECT \r\n a1, " \"a bc \t" as a2\r FROM c\rWHERE ( a3 = 5 AND a2>1 )`
	dst = `SELECT a1," \"a bc \t"as a2 FROM c WHERE(a3=5 AND a2>1)`
	TringSentence(src, dst)

}

func TringSentence(src string, dst string) {
	if dst == *trim_sentence(src) {
		fmt.Println("True!")
	}
}

// 返回值说明：go语言返回字符串时默认是返回整个字符串而不是它的地址，需要显示声明指针，可以显著降低时间空间成本
func trim_sentence(src string) *string {
	flag := false
	s := make([]byte, 0)
	l := len(src)
LABEL:
	for i := 0; i < l; i++ {
		// 判断引号，以及是否成对
		if src[i] == '"' && i < l-1 && src[i-1] != '\\' {
			for j := i + 1; j < l; j++ {
				if src[j] == '"' && src[j-1] != '\\' {
					flag = true
					s = append(s, src[i])
					continue LABEL
				}
			}
		}
		// 有引号时无差别复制
		if flag {

			s = append(s, src[i])
			// 当引号闭合继续进行循环
			if src[i] == '"' && src[i-1] != '\\' {
				flag = false
				continue LABEL
			}
			// 没有引号时进行条件选择复制
		} else {
			// 判断 \ 后字符 决定是否加入dst
			if src[i] == '\\' {
				switch src[i+1] {
				case 'r', 't', 'n', 'v', 'f':
					{
						i++
						if unicode.IsLetter(rune(src[i+1])) {
							s = append(s, ' ')
						}
						continue
					}
				default:
					{
						s = append(s, src[i])
					}
				}
				// 判断 空格 后字符 决定是否加入dst
			}
			if unicode.IsSpace(rune(src[i])) && i <= l-1 {
				switch {
				case (unicode.IsNumber(rune(src[i-1])) || unicode.IsLetter(rune(src[i-1]))) && (unicode.IsLetter(rune(src[i+1])) || unicode.IsNumber(rune(src[i+1]))):
					s = append(s, src[i])
					continue LABEL
				default:
					continue LABEL

				}
			}
			s = append(s, src[i])
		}
	}
	dst := string(s)
	return &dst
}
