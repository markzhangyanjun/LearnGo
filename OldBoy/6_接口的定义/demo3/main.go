package main


//a. Golang中的接口，不需要显示的实现。只要一个变量，含有接口类型中的所有方法，那么这个变量就实现这个接口。因此，golang中没有implement 类似的关键字
//b. 如果一个变量含有了多个interface类型的方法，那么这个变量就实现了多个接口。
//c. 如果一个变量只含有了1个interface的方部分方法，那么这个变量没有实现 这个接口。
import "fmt"

type Test1 interface {
	Print()
	Sleep()
}

type Test2 interface {
	Run()
}

type Student struct{
	Name string
	Age  int
	Score int
}

func(p *Student)Print(){
	fmt.Printf("name=%s,age=%d,score=%d\n",p.Name,p.Age,p.Score)
}
func(p *Student)Sleep(){
	fmt.Println("stu is sleep")
}

func(p *Student)Run(){
	fmt.Println("stu is running")
}



func main() {
	var t1 Test1
	var s Student
	s = Student{"mark",18,98}

	t1=&s
	t1.Print()

	var t2 Test2
	t2 = &s
	t2.Run()

}
