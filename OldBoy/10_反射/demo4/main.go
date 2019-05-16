package main

import (
	"reflect"
	"fmt"
)

//获取变量的值
//reflect.ValueOf(x).Float()
//reflect.ValueOf(x).Int()
//reflect.ValueOf(x).String()
//reflect.ValueOf(x).Bool()
func Test(p interface{}){
	fmt.Printf("%T\n",p)

	v :=reflect.ValueOf(p)
	fmt.Println(v)
	fmt.Printf("%T\n",v)
	res :=v.Int()
	fmt.Println(res)
	fmt.Printf("%T\n",res)



}
func main() {

    var a int = 10

    Test(a)

}
