package main

import (
	"fmt"
	"sync"
)

func add1(a *int,wg sync.WaitGroup){
	defer wg.Done()
	fmt.Println("enter add1")
	wg.Add(1)
	for i:=0;i<=1000;i++{
		*a = *a + i

	}
	fmt.Println("add1 end")

}

func add2(a *int,wg sync.WaitGroup){
	defer wg.Done()
	fmt.Println("enter add2")
	wg.Add(1)
	for i:=0;i<=1000;i++{
		*a = *a + i

	}
	fmt.Println("add2 end")

}

func main() {

	var wg sync.WaitGroup

	var a int =0
	go add1(&a,wg)
	go add2(&a,wg)




	wg.Wait()
	fmt.Println("main exist")


}
