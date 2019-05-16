package __判断一个变量是否是指定接口

import "fmt"

type Reader1 interface {
	Read()
}

type Writer1 interface {
	Write()
}

type ReadWriter1 interface {
	Reader1
	Writer1
}

type File1 struct{

}

func(f *File1)Read1(){

	fmt.Println("readFile")

}
func(f *File1)Write1(){
	fmt.Println("writeFile")


}

func myTest1(rw ReadWriter1){
	rw.Read()
	rw.Write()

}
func main() {

	//var f *File1
	//
	//var b interface{}
	//b = f
	//if v,ok:=b.(ReadWriter1);ok{
	//	fmt.Println(v,ok)

	//}




}


