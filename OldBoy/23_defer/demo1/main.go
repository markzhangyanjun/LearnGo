package main

import "fmt"

func main() {
	var i int = 0

	defer fmt.Println(i)
	defer fmt.Println("second")

	i =10
	fmt.Println(i)
	defer fmt.Println(i)



}

