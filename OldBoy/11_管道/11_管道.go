package _1_管道

import "fmt"

func main() {

	pipe := make(chan int, 3)

	pipe <- 1
	pipe <- 2
	pipe <- 3


	t1 := <- pipe
	fmt.Println("t1=",t1)
	t2 := <- pipe
	fmt.Println("t2=",t2)

	t3:= <- pipe
	fmt.Println("t3=",t3)

	t4 := <- pipe
	fmt.Println("t4=",t4)

}