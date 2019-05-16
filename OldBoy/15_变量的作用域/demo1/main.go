package demo1

import "fmt"

func test(){
	var a =100
	fmt.Println("a=",a)
    var b = 100
	for i:=0;i<100;i ++{
		b =i*2
		fmt.Println("b=",b)
	}
	fmt.Println("b=",b)
}


var a = "G"

func n(){
	fmt.Println(a)
}
func m(){
	a="O"
	fmt.Println(a)
}



func main() {

	n()
	m()
	n()

}
