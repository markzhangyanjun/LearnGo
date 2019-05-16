package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct{
	Name string
	Id   int
	Age  int
}



type StudentArray []Student

func(p StudentArray)Len()int{
	fmt.Println(p)
	return len(p)
}

func(p StudentArray)Swap(i,j int){
	fmt.Println(p)
	p[i] ,p[j] = p[j],p[i]
}
//根据姓名排序
type SortByName struct{
	StudentArray
}

func(p SortByName)Less(i,j int)bool{
	return p.StudentArray[i].Name>p.StudentArray[j].Name

}

//根据ID排序
type SortByID struct{
	StudentArray
}

func(p SortByID)Less(i,j int)bool{
	return p.StudentArray[i].Id>p.StudentArray[j].Id
}





func main() {

	var stus []Student

	for i:=1;i<5;i++{

		stu:=Student{
			Name :fmt.Sprintf("sut%d",rand.Intn(100)),
			Id :i*10,
			Age:rand.Intn(100),

		}
		stus=append(stus,stu)
	}
	sbn := SortByName{stus}
	sort.Sort(sbn)
	for _,value :=range stus{
		fmt.Println(value)
	}





	fmt.Println("############")
	sbID:=SortByID{stus}
    r:= sbID.Len()
    fmt.Println("r=",r)
	sort.Sort(sbID)
	for _,v :=range stus{
		fmt.Println(v)
	}



}
