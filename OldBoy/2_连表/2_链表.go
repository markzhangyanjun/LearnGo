package main

import (
	"fmt"
	//"math/rand"
)

type Studenta struct {
	Name string
	Age  int
	Score float32
	next  *Studenta
}



func main(){
	var head Studenta

	head.Name = "zyj"
	head.Age = 10
	head.Score =100

	var head1 Studenta
	head1.Name = "mark"
	head1.Age = 10
	head1.Score =100

	var head2 Studenta

	head2.Name = "jjj"
	head2.Age = 18
	head2.Score = 22


	head.next = &head1

	head1.next = &head2

    var p *Studenta = &head
	for p !=nil{
		fmt.Println(*p)
		p = p.next

	}


    //尾部插入
    //var tail =&head
    //for i :=0;i<10;i++{
	//
    //	var stu =Studenta{
    //		Name : fmt.Sprintf("stu%d",i),
    //		Age :rand.Intn(100),
    //		Score :rand.Float32()*100,
	//	}
	//
    //	tail.next =&stu
    //	tail=&stu
	//
	//}
	//
	//var p *Studenta = &head
	//for p !=nil{
	//	fmt.Println(*p)
	//	p = p.next
	//
	//}


	//头部插入

	//for i :=0;i<10;i++{
	//	var stu = Studenta{
	//		Name : fmt.Sprintf("stu%d",i),
	//		Age :rand.Intn(100),
	//		Score :rand.Float32()*100,
	//
	//	}
	//	stu.next = &head
	//	head = stu
	//	fmt.Println(stu)
	//
	//
	//}

}



