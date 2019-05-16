package ___和_



//& 是取地址符号 , 即取得某个变量的地址 , 如 ; &a
//*是指针运算符 , 可以表示一个变量是指针类型 , 也可以表示一个指针变量所指向的存储单元 , 也就是这个地址所存储的值 .
import "fmt"

func main() {

	 a := "hello"

	 b := &a
	 fmt.Println(b)

	 *b = "word"
	 fmt.Println(a)


}