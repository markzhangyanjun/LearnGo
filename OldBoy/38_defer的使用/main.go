package main

import "fmt"
/*
1. 当函数返回时，执行defer语句。因此，可以用来做资源清理
2. 多个defer语句，按先进后出的方式执行
3. defer语句中的变量，在defer声明时就决定了。
*/

func main() {

   //defer语句中的变量，在defer声明时就决定了。
   var i int= 0
   defer fmt.Println(i)
	//多个defer语句，按先进后出的方式执行
   defer fmt.Println("this is test")


   i = 10
   fmt.Println(i)


   //例子：
   for i:=0;i<5;i++{
   	  defer fmt.Printf("%d\n",i)
	}


}
