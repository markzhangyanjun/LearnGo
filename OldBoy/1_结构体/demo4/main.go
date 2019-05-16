package main

import (
	"fmt"
	"time"
)

//结构体中字段可以没有名字，即匿名字段

type Car struct{
	name  string
	age  int
}

type Train struct{
	Car
	int
	start  time.Time
}
func main() {
    var t Train
    t.name = "train"
    t.age = 18
    t.int = 200
    fmt.Println(t)

    //也可以这么写
    t.Car.name = "train001"
    t.Car.age = 100
    t.int = 1000


}
