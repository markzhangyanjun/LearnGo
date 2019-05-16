package main

import (
	"fmt"
	"sort"
)

type Intger []int

func(p Intger)Len()int{
	return len(p)
}
func(p Intger)Less(i,j int)bool{
	return p[i] < p[j]
}
func(p Intger)Swap(i,j int){
	p[i],p[j] = p[j],p[i]
}

func main() {

	i := Intger{1,5,3,8,2,7,4}
	sort.Sort(i)
	fmt.Println(i)

}
