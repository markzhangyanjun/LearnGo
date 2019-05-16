package main

import "fmt"
func main() {
	s := []int{1,2,3}
	s1 := []int{4,5,6}
	s = append(s,s1...)
	fmt.Println(s)
}
