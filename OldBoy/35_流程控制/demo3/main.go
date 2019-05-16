package main

import (
	"fmt"
	"strings"
)
func main() {
	str3:="A"
	for i :=1;i<10;i++{

		//重复count次str
		result:=strings.Repeat(str3, i)
		result += "\n"
		fmt.Println(result)
	}
}
