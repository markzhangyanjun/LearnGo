package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//随机生成一个1～100的整数
	 r := rand.Intn(100)
	 fmt.Println(r)
	 for{

		 //接受用户输入数据
		 var d int
		 fmt.Scanf("%d\n",&d)

		 switch {
		 case d==r:
			 fmt.Println("OK")
		 case d > r:
		 	fmt.Println("大了")
		 case d < r:
			 fmt.Println("小了")


		 }

	 }


}
