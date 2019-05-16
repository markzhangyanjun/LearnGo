package main

import "fmt"

func main() {
	for i:=0;i< 5;i++{
		defer fmt.Printf("defer[%d]\n",i)
	}
}
