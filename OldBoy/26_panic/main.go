package main

import "fmt"

func test(){

	defer func(){
		if err:= recover();err != nil{
			fmt.Println(err)
		}
	}()

	b :=0
	a := 100/b
	fmt.Println(a)
}
func main() {
    for{
		test()
	}

}
