package main

import (
	"fmt"
	"sort"
)
//数组排序
func testIntSort(){
	var a =[...]int{1,4,8,2,5}
	sort.Ints(a[:])
	fmt.Println(a)
}

func testStringSort(){
	a := []string{"a","c","b","w","m"}
	sort.Strings(a)
	fmt.Println(a)
}

func testFloatSort(){
	a := []float64{2.5,3.8,4.2,1.5,3,0.4,0.04}
	sort.Float64s(a)
	fmt.Println(a)
}

func testIntSearch(){
	var a =[...]int{1,4,8,2,5}
	sort.Ints(a[:])
	fmt.Println(a)
	index :=sort.SearchInts(a[:],2)
	fmt.Println(index)

}

func main() {

	testIntSort()
	testStringSort()
	testFloatSort()
	testIntSearch()

}
