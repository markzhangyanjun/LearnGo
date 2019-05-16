package main

import "fmt"
type Student struct{
	Name  string
	Age   int
}

func Test(a interface{}){

	b,ok :=a.(Student)
	if ok ==false{
		fmt.Println("convert failed")
	}
	fmt.Println(b)

}

func main() {

	var a Student
	Test(a)


}
