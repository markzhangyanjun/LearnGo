package main

import "fmt"

//Golang中的方法是作用在特定类型的变量上，因此自定义类型，都可以有方法，而不仅仅是struct
//定义：func (recevier type) methodName(参数列表)(返回值列表){}


type  integer int

func(p integer)print(){
	fmt.Println(p)
}

func(p *integer)set(b integer){
	*p = b

}



type Student struct{
	Name string   `json:"name"`
	Age int     `json:"age"`
	Score int   `json:"score"`
}

func(p *Student)init(name string ,age int,score int){

	p.Name = name
	p.Age = age
	p.Score = score
}

func(p Student)get()Student{
	return p
}

func main() {
     //var stu  Student
     //stu.init("stu",18,29)
	 //
     //fmt.Println(stu)
     //stu1 := stu.get()
     //fmt.Println(stu1)

     var i integer=1
     i.print()

     i.set(100)
     fmt.Println(i)
}
