package main
import "fmt"
//动态的获取变量的类型


type Student struct{
	Name string
	Age  int
}

func GetParamsType(items ...interface{}){

	for _,value :=range items{
		switch value.(type){
		case bool:
			fmt.Printf("%v type is  bool\n",value)
		case int:
			fmt.Printf("%v type is  int\n",value)
		case string:
			fmt.Printf("%v type is  string\n",value)
		case Student:
			fmt.Printf("%v type is  Student\n",value)
		case *Student:
			fmt.Printf("%v type is  *Student\n",value)

		}

	}

}

func main() {
    s := Student{"abc",18}

	GetParamsType(10,true,"abc",s,&s)

}
