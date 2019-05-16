package main

import (
	"fmt"
	sum "LearnGo/OldBoy/29_给包起别名/add" //起别名
)

func main() {
	a:= 1
	b:= 1

	c :=sum.Add(a,b)
	fmt.Println(c)
}