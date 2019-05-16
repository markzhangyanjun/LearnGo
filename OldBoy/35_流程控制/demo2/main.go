package main

import "math/rand"
import "fmt"

/*
猜数字，写一个程序，随机生成一个0到100的整数n，然后用户在终端，
输入数字，如果和n相等，则提示用户猜对了。如果不相等，则提示用户，大于或小于n。
*/
func main() {

	//生成随机整数
	n := rand.Intn(100 )
	fmt.Println("n=",n)

	for{

		var num int
		fmt.Scanf("%d\n",&num)
		var flag bool=false

		switch{
		case n==num:
			fmt.Println("ok")
		    flag =true
		case n > num:
			fmt.Println("less")
		case n < num:
			fmt.Println("bigger")
		}
		if flag == true{

			break
		}
	}


}
