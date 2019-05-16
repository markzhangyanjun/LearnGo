package main

//Interface类型可以定义一组方法，但是这些不需要实现。并且interface不能
//包含任何变量。

import "fmt"

type Test interface {
	Hello()

}


type Carer interface {
	GetName() string
	Run()
	DiDi()
}

type BMW struct{
	BName string
	LName string
}
type BYD struct{
	Name string
}


func(p *BYD)GetName()(bigName  string ){
	bigName = p.Name
	return
}

func (p *BYD)Run(){
	fmt.Println("BYD is Running ")
}

func(p *BYD)DiDi(){
	fmt.Println("BYD is DIDIing")
}

func(p *BMW)GetName()(bigName  string ){
	bigName = p.BName
	return
}

func (p *BMW)Run(){
	fmt.Println("BMW is Running ")
}

func(p *BMW)DiDi(){
	fmt.Println("BMW is DIDIing")
}

func(p *BMW)Hello(){
	fmt.Println("hello")
}



func main() {

	var car Carer
	var bmw BMW
	var byd BYD
	//bmw.BName="BWM"
	//bmw.LName = "bmw"
	car = &bmw
	car.Run()



	car = &byd
	car.Run()







}
