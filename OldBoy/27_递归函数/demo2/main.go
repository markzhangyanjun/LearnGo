package main

import "fmt"
//阶乘
func calc(n int)int{
	if n==1{
		return 1
	}
	num:=calc(n-1)*n
	fmt.Println(num)
	return calc(n-1)*n

}

func main() {

	num := calc(5)
	fmt.Println(num)

}
