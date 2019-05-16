package dem2

import (
	"fmt"
	"strings"
)

func urlProcess(url string)string{

	//判断是以xxx开头
	result := strings.HasPrefix(url,"http://")


	if !result{
		url = fmt.Sprintf("http://%s",url)

	}

	return url

}

func pathProcess(path string)(r_path string){
	//判断以XXX结尾
	result :=strings.HasSuffix(path,"/")
	fmt.Println(result)
	if !result{
		path =fmt.Sprintf("%s/",path)
	}

	return path

}

func main() {
	var (
		url string
		path string
	)

	fmt.Scanf("%s%s",&url,&path)

	url =urlProcess(url)
	fmt.Println(url)

	path = pathProcess(path)
	fmt.Println(path)

}
