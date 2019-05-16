package main
import "fmt"
//判断一个变量是否实现了指定的接口
type Read interface {
	Read()
}

type Student struct{

}

//func(p *Student)Read(){
//	fmt.Println("read book")
//}
func main() {

	var s *Student

    var b interface{}
	b = s
	if sx,ok :=b.(Read);ok{

		fmt.Println("实现",ok)
		fmt.Println(sx)
	}else{
		fmt.Println(ok)
		fmt.Println(sx)
	}


}
