package main

import "fmt"
//new：用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
// make：用来分配内存，主要用来分配引用类型，比如chan、map、slice
func main() {
	//new 返回的是地址
	j := new(int)
	*j= 10
	fmt.Println(j)
	fmt.Println(*j)

	str := new(string)
	*str = "helo"


	s1 := new([]int)
	fmt.Println(s1)
	s2:=make([]int,0)
	s2[0]=100
	fmt.Println(s2)


}
