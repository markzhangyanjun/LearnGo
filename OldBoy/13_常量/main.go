package main

import (
	"fmt"
	"time"
)

const (
	Male = 1
	Female = 2
)
func main(){
	for{
		second := time.Now().Unix()
		fmt.Println(second)
		if (second % Female == 0){
			fmt.Println("female")
			time.Sleep(5*time.Second)
		} else{
			fmt.Println("man")
			time.Sleep(5*time.Second)
		}
	}

}
