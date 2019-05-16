package main

import "fmt"

func swap(a, b *int)(){
	tmp := *a
	*a = *b
	*b = tmp
}

func main(){

	var a int = 1
	var b int = 10
	swap(&a,&b)
	fmt.Println("a=",a)
	fmt.Println("b=",b)

}
