package test

import "fmt"



var Name string = "xxxxx"
var Age  int = 18
func init(){

	fmt.Println("this is test init")
	fmt.Println("this is test package Name=",Name)
	fmt.Println( "this is test package Age=",Age)

}
