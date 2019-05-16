package main

import (
	"fmt"
	"strconv"
)


//写一个程序，从终端读取输入，并转成整数，如果转成整数出错，
func main() {

	//从终端读取数据
    var str1 string
	fmt.Scanf("%s",&str1)
    //将字符串转换为整数
	a,err := strconv.Atoi(str1)
	if err !=nil{
		fmt.Println("can not convert to int")

	}else{
		fmt.Println("a=",a)
	}


}
