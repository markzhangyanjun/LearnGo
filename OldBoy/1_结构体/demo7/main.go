package main

import "fmt"

//多重继承
//如果一个struct嵌套了多个匿名结构体，那么这个结构可以直接访问
//多个匿名结构体的方法，从而实现了多重继承。

type Person struct{
	Name string
	Age  int
}

func(p *Person)init(name string ,age int){
	p.Name = name
	p.Age = age
}


type Father struct{
	Work   string
}

func(p *Father)work(work string){
	p.Work=work

}


type Son struct{
	Person
	Father
}

func main() {
	var s Son
	//s.Name = "小明"
	//s.Age = 18
	//s.Work = "程序员"
	//fmt.Printf("s=%+v\n",s)

	s.init("小红",18)
	fmt.Printf("s=%+v\n",s)
	s.work("工人")
	fmt.Printf("s=%+v\n",s)

}


