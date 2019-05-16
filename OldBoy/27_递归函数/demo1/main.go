package main

import (
	"fmt"
	"time"
)

func recusive(n int){
	fmt.Printf("start,n=%d\n",n)
	time.Sleep(time.Second)
	if n > 10{
		fmt.Printf("end, n=%d\n",n)
		return
	}
	recusive(n+1)
}
func main() {
	recusive(0)
}
