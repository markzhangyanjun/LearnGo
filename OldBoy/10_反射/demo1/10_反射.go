package demo1
//反射：可以在运行时动态获取变量的相关信息

import (
	"fmt"
	"reflect"
)



type SS struct{
	Name string
	Age  int
	Score float32

}
func test(b interface{}){

	t :=reflect.TypeOf(b) //类型
	fmt.Println(t)

	v := reflect.ValueOf(b)
	fmt.Println(v)

	k :=v.Kind() //类别
	fmt.Println(k)

	iv :=v.Interface()
	fmt.Println(iv)
    //类型断言
	stu,ok :=iv.(SS)
	if ok{
      fmt.Printf("%v %T\n",stu,stu)
	}
}

func testInt(b interface{}){
	va1 := reflect.ValueOf(b)
	fmt.Println(va1)
	c := va1.Int()
	fmt.Printf("%d,%T\n",c,c)

}
//反射
//获取便量值
//reflect.ValueOf(x).Float()
//reflect.ValueOf(x).Int()
//reflect.ValueOf(x).String()
//reflect.ValueOf(x).Bool()

func main() {

	//var a int =200
	//
	//test(a)

	//var  b SS = SS{
	//	Name:"maek",
	//	Age :18,
	//	Score :99,
	//}
	//test(b)

	testInt(123)

}
