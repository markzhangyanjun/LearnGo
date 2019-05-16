package add

import "fmt"

import _ "LearnGo/OldBoy/12_变量的使用/test"



func init(){
	fmt.Println("this is add init")
	Name = "mark"
	Age = 18

}

var Name string = "xxxxx"
var Age  int = 18
