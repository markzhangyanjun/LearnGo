package main

import "fmt"

type integer int

type  Student struct{
	Name  string
	Age   int
}

type Stu   Student  //定义类型相当于取别名，想要与原类型进行赋值需要进行强制转换

func main() {
	var i integer =1000
	fmt.Println(i)

	var j int =100
	fmt.Println(j)
	//j = i 此处保存，因为类型不一样

	var s1 Student=Student{"zyj",18}


	var s2 Stu =Stu{"mark",18}


	//s1 = s2  此处报错因为类型不同 ,想要赋值可以进行强制转换
	s1 = Student(s2)
	fmt.Println(s1)

	var s3 Student=Student{"Mark",18}
	fmt.Println(s3)

	s3 = s1
	fmt.Println(s3)

}
