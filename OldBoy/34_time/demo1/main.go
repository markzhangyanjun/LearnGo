package main

import (
	"time"
	"fmt"
	)


/*
1. time.Time类型，用来表示时间
2. 获取当前时间， now := time.Now()
3. time.Now().Day()，time.Now().Minute()，time.Now().Month()，time.Now().Year()
4. 格式化，fmt.Printf(“%02d/%02d%02d %02d:%02d:%02d”, now.Year()…)
*/



func main() {

	t1 := time.Now()
	fmt.Println("t1=",t1)

	t2:= time.Now().Day()
	fmt.Println("t2=",t2)

	t3:= time.Now().Minute()
	fmt.Println("t3=",t3)

	t4:= time.Now().Month()
	fmt.Println("t4=",t4)

	t5:= time.Now().Year()
	fmt.Println("t5=",t5)

	now := time.Now()

	fmt.Println(now.Format("02/1/2006 15:04"))
	fmt.Println(now.Format("2006/1/02 15:04"))
	fmt.Println(now.Format("2006/1/02"))



}
