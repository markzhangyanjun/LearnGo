package init
import (
	"fmt"
	_"LearnGo/OldBoy/30_Init函数/test" //导入一个包用于初始化,里面的变量没有应用
	)


var Name string = "xxxx"
var Age  int =10

func init(){
	 fmt.Println("hello")
	 Name = "mark"
	 Age =18
}
