package app

import "fmt"

func PrintHello() {
	// 切片操作
	s1 := "hello "
	s2 := "wworld"
	s3 := s1 + s2[1:]

	// 自定义输出
	s4 := `eret
	       ewfwegfwe `
	fmt.Println(s3, s4)
}
