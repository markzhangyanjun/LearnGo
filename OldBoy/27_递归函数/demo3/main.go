package main

import "fmt"

//斐波那契数

func calc(n int)int{
	if n <= 1{
		return 1
	}
	return calc(n-1) +calc(n-2)
}

func main() {

   num := calc(4  )
   fmt.Println(num)


}
