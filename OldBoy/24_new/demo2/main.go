package main

import "fmt"
//简单的说，new只分配内存，make用于slice，map，和channel的初始化。
func test(){
	s1 := new([]int)
	fmt.Println(s1)
	s2:=make([]int,10)
	fmt.Println(s2)

	(*s1)[0] =100
	s2[0]=100

	return
}

func main() {

	test()

}