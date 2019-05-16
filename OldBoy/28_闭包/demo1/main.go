package main
import "fmt"

// 闭包：一个函数和与其相关的引用环境组合而成的实体

func Add()func(int) int{
	var x int
	return func(d int)int{
		x += d
		return x
	}
}



func main() {

	f:=Add()
	fmt.Println(f(1))
	fmt.Println(f(100))


}
