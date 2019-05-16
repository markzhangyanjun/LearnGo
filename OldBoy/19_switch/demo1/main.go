package demo1

import "fmt"

func main() {
	var a int = 0

	switch a {
	case 0:
		fmt.Println("a=",0)
	    //fallthrough ,可以让下面的一行语句也执行
	    fallthrough
	case 1:
		fmt.Println("a=",1)
	    //fallthrough
	case 2:

		fmt.Println("a=",2)

	}



}
