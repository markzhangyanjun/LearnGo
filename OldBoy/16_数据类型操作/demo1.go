package main

import "fmt"

func main() {

	var a int
	var b int32
	a = 15
	b = int32(a + a ) //强制类型转换
	fmt.Println(b)

	var c float32
	c = 10.333333333333333333
	fmt.Println(c)
	d := float64(c)
	fmt.Println(d)


}
