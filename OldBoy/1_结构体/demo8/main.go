package main

import "fmt"

//如果一个变量实现了String()这个方法，那么fmt.Println默认会调用这个
//变量的String()进行输出。

type Student struct{
	Name string
	Age  int
}

func(p *Student)String()(str string){
	str =fmt.Sprintf("name=%s,age=%d",p.Name,p.Age)
	return str
}

func main() {
	s := Student{"zyj",1 }
	fmt.Printf("%v",&s)
}
