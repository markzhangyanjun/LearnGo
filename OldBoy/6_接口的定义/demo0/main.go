package main

import "fmt"

//空接口
//空接口没有任何方法，所以所有类型都实现了空接口。

func main() {
	var a int
	var b interface{}
	b =a
	fmt.Println(b)

	fmt.Printf("%T",b)
}

