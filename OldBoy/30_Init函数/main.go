package main

import (
	a "LearnGo/OldBoy/30_Init函数/init" //导入一个包做初始化但是应用了里面的变量不需要加下划线
	"fmt"
	)

func main() {
     fmt.Println(a.Name)
     fmt.Println(a.Age)

}
