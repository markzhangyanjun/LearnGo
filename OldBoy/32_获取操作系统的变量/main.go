package main
//写一个程序获取当前运行的操作系统名称和PATH环境环境变量的值，并打印在终端
import(
	"os"
	"fmt"
)


func main() {

    gopath :=os.Getenv("HOME")
    fmt.Println(gopath)


}
