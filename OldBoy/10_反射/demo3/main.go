package main

import (
	"reflect"
	"fmt"
)
type Student struct {
	Name string
	Age  int
}

func Test(a interface{}){
	t :=reflect.TypeOf(a)
	fmt.Println(t)

	v :=reflect.ValueOf(a)
	fmt.Println(v)

	k :=v.Kind()
	fmt.Println(k)

	iv := v.Interface()  //将Student转换成interface类型
	fmt.Println(iv)
	fmt.Printf("%T\n",iv)

	//将interface转换为原来的类型
	value,ok :=iv.(Student)
	if ok ==false{
		fmt.Println("类型转换失败")
		return
	}
	fmt.Printf("value type %T",value)


}


func main() {

	var a Student =Student{"mark",18}

	Test(a)


}
