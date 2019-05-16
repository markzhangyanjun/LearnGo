package main

import (
	"fmt"
	"os"
)

//自定义错误

type Patherror struct{
	path string
	op   string
	message string

}

func (p *Patherror)Error()string{

	return fmt.Sprintf("path=%s op=%s message=%s",p.path,p.op,p.message)

}

func Open(filename string)error{
	file,err :=os.Open("test.sls")
	if err !=nil{
		fmt.Printf("%T\n",err.Error())
		fmt.Println(err.Error())
		return &Patherror{ "test.sls", "read", err.Error()}
	}

	defer file.Close()
	return nil

}

func main() {

	res := Open("hello")
	fmt.Println(res)

}
