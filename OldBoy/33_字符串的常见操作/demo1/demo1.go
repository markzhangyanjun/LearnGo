package main

import (
	"strings"
	"fmt"
)

/*
strings.Replace(str string, old string, new string, n int)：字符串替换
strings.Count(str string, substr string)int：字符串计数
strings.Repeat(str string, count int)string：重复count次str
strings.ToLower(str string)string：转为小写
strings.ToUpper(str string)string：转为大写

strings.TrimSpace(str string)：去掉字符串首尾空白字符
strings.Trim(str string, cut string)：去掉字符串首尾cut字符
strings.TrimLeft(str string, cut string)：去掉字符串首cut字符
strings.TrimRight(str string, cut string)：去掉字符串首cut字符


strings.Field(str string)：返回str空格分隔的所有子串的slice
strings.Split(str string, split string)：返回str split分隔的所有子串的slice
strings.Join(s1 []string, sep string)：用sep把s1中的所有元素链接起来

 strings.Itoa(i int)：把一个整数i转成字符串
strings.Atoi(str string)(int, error)：把一个字符串转成整数
*/

func main() {
    //字符串替换
	str1 := "helloworldhellobaby"
	str1 =strings.Replace(str1,"hello", "@", 2)
	fmt.Println(str1)

	//字符串计数
	str2:= "asdfghjklqwertyuiopzxcvbnmasdfghjkl"
	count := strings.Count(str2, "a")
	fmt.Println("count=",count)

	//重复count次str
	str3:="hello python "
	result :=strings.Repeat(str3, 2)
	fmt.Println("result=",result)

	//转为小写
	str4:="HelLoWorld"
	result= strings.ToLower(str4)
	fmt.Println("result=",result)

	//转为大写
	str5:="hello python"
	result= strings.ToUpper(str5)
	fmt.Println("result=",result)

	//去掉字符串首尾空白字符
	str6:="   hello"
	result =strings.TrimSpace(str6)
	fmt.Println("result=",result)

	//去掉字符串首尾cut字符
	str7 := "@@@hello python@@@"
	result = strings.Trim(str7, "@")
	fmt.Println("result=",result)

	//去掉字符串首cut字符
	str8:="@@@hello python@@@"
	result = strings.TrimLeft(str8, "@")
	fmt.Println("result=",result)

	//去掉字符串首cut字符
	str9:= "@@@hello python@@@"
	result = strings.TrimRight(str9, "@")
	fmt.Println("result=",result)

	//返回str空格分隔的所有子串的slice
	str11:= "  hello python golang  "
	result_slice := strings.Fields(str11)
	fmt.Println(result_slice)

	//返回str split分隔的所有子串的slice
	str12:= "hello@python@golang"
	result_slice = strings.Split(str12, "@")
	fmt.Println(result_slice)

	//用sep把s1中的所有元素链接起来

	ss:=[]string{"hello","python","golang"}
	buf:= strings.Join(ss, "-")
	fmt.Println("buf=",buf)



























}
