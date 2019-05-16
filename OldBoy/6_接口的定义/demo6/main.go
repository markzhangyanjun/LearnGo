package main

import "fmt"

//接口嵌套


type Read interface {
	Read()
}
type Write interface {
	Write()
}

type ReadWrite interface {
	Read
	Write
}
type File struct{


}

func(p *File)Write(){
	fmt.Println("read data")
}

func(p *File)Read(){
	fmt.Println("write data")
}

func Test(rw ReadWrite){

	rw.Read()
	rw.Write()

}
func main() {

	var f File
	Test(&f)
	var rw ReadWrite
	rw=&f
	rw.Write()


}
