package main

import "fmt"

func main() {
	str:="我是中国人"

	for index,value :=range str{
		fmt.Println(index,value)
	}

	s :=[]rune(str)
	fmt.Println(len(s))
	s[0]='他'
	fmt.Println(string(s))



	//str1:= "abc"
	//for index,value :=range str1{
	//	fmt.Println(index,value)
	//}
	//
	//str2:= "abc"
	//fmt.Println([]byte(str2))
	//fmt.Println(string([]byte(str2)))
}



