package __匿名字段

import (
	"fmt"
)

type Car struct{
	Name string
	Age  int
}

type Train struct{
	Car
	int
}


//& 是取地址符号 , 即取得某个变量的地址 , 如 ; &a
//*是指针运算符 , 可以表示一个变量是指针类型 , 也可以表示一个指针变量所指向的存储单元 , 也就是这个地址所存储的值 .
func(t *Train)Run(){

	t.int = 300


}

func main() {
	var train Train
	train.int = 100
	train.Name = "zyj"
	train.Age = 18
    train.Run()
	fmt.Println(train)



}