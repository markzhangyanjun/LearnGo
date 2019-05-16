package demo2

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct{

}

func(f *File)Read(){

	fmt.Println("readFile")

}
func(f *File)Write(){
	fmt.Println("writeFile")


}

func myTest(rw ReadWriter){
	rw.Read()
	rw.Write()

}
func main() {

	var f File

	myTest(&f)


}
