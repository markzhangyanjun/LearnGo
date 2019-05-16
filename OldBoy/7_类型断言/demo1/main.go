package main
//空接口的使用
//如果一个变量含有了多个interface类型的方法，那么这个变量就实现了多个接口。如果这个接口中没有任何方法
//那么任何类型都可以给该接口赋值
import "fmt"

func main() {
	var a interface{}

	var b int
	a = b
	c,ok:= a.(int)
	if ok ==false{
		fmt.Println("类型转换失败")
		return
	}
	d :=c+3
	fmt.Println("d=",d)
	fmt.Printf("%T\n",c)
	fmt.Printf("%T",a)

}
