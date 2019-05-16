package main

import "fmt"

//继承
//如果一个struct嵌套了另一个匿名结构体，那么这个结构可以直接访问
//匿名结构体的方法，从而实现了继承。


type Car struct{
	weight int
	name  string
}
func(p *Car)Run(){
	fmt.Println("is running")
}
type Bike struct{
	Car
	lunzi int
}

type Train struct{
	Car
	lunzi int
}

func main() {
	var b Bike
	b.weight = 100
	b.name = "永久"
	b.lunzi = 2
	fmt.Printf("b=%+v\n",b)
	b.Run()

	var t Train
	t.weight = 2000
	t.name = "火车"
	t.lunzi=100
	fmt.Printf("t=%+v\n",t)
	t.Run()

}

