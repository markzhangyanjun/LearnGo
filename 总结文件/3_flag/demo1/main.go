package main

import (
"flag"
"fmt"
)

func main() {

	var num int
	var mode string
	var encrypt bool

	// usage
	// cmd -flag
	// cmd --flag
	// cmd -flag=
	// cmd --flag=

	// 第一步，定义Flag参数
	flag.IntVar(&num, "num", 16, "-num the password length")
	//var num = flag.Int("num", 16, "-num the password length")


	flag.StringVar(&mode, "mode", "mix", "-mode the password generate mode")
	// var mode = flag.String("mode", "mix", "-mode the password generate mode")

	flag.BoolVar(&encrypt, "encrypt",true, "-encrypt the password use RSA")
	//var encrypt = flag.Bool("encrypt",true, "-encrypt the password use RSA")


	// 第二步，调用flag.Parse()解析命令行参数到定义的Flag
	flag.Parse()

	// 调用Parse解析后，就可以直接使用绑定的变量了
	fmt.Printf("num:%d mode:%s encrypt:%v\n", num, mode, encrypt)

}