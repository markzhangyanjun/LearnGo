package main

import "time"
import "fmt"

//写一个程序，获取当前时间，并格式化成 2017/06/15 08:05:00形式
func main() {

	now := time.Now()
	//必须是这个格式
	fmt.Println(now.Format("2006/01/02 03:04:05"))
	fmt.Println(now.Format("2006-01-02 03:04:05"))

}
