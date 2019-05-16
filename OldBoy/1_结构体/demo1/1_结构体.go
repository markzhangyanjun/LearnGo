package demo1

import "fmt"

type Student struct{
	Name string
	Age  int
	Score float32
}


func main(){
	var stu Student
	stu.Name = "mark"
	stu.Age = 18
	stu.Score = 100

	fmt.Println("stu=",stu)
	fmt.Printf("Name=%p,Age=%p,Score=%p\n",&stu.Name,&stu.Age,&stu.Score)


	//初始化
	var stu1  *Student = &Student{Name:"zyj",Age:18,Score:20}
	fmt.Println("stu1=",stu1)

	//打印地址
	fmt.Printf("%p\n",&stu)
	fmt.Printf("Name=%p\n",&stu.Name)
	fmt.Printf("Age=%p",&stu.Age)
}


