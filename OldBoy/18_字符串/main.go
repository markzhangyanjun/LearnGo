package main

import "fmt"

func main() {
    //``:是什么样就打印成什么样
	var str = "hello world\n"
	var str1 = `
窗前明月光\n
疑是地上霜,
举头望明月,
低头思故乡
`
	fmt.Println(str)
	fmt.Println(str1)
    //定义字符用单引号
	var b byte = 'b'
	fmt.Println(b)
	fmt.Printf("%c\n",b)


    var str2 = "hello"
    var str3 = "world"

    str4 := str2 + str3
    fmt.Println(str4)
    str5 := fmt.Sprintf("%s",str4)
	fmt.Println(str5)



}
