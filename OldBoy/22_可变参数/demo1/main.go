package main

import "fmt"

func add(arg ...int)int{
	var sum int = 0
	for i:=0;i<len(arg);i++{
		sum =arg[i] +sum
	}

	return sum

}

//字符串拼接
func addStr(str string,arg ...string)string{
	str += " "
	for i:=0;i<len(arg);i++{
		str += arg[i]
	}

	return str
}

func main() {


	//sum:=add(1,2,3,4,5)
	//fmt.Println("sum=",sum)
	//
	//sum1 := add()
	//fmt.Println("sum1=",sum1)

	var str string = "hello"
	str1:= addStr(str,"w","o","r","l","d")
	fmt.Println(str1)

}



