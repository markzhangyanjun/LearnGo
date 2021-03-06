package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main() {


	config := make(map[string]interface{})
	config["filename"] = "/Users/zhangyanjun/Desktop/GO/src/LearnGo/OldBoy/test_log/logs/logcollect.log"
	config["level"] = logs.LevelDebug

	configStr, err := json.Marshal(config)
	if err != nil {
	fmt.Println("marshal failed,err:", err)
	return
	}
	fmt.Println(string(configStr))
	logs.SetLogger(logs.AdapterFile, string(configStr))
	logs.SetLogger(logs.AdapterConsole)
	logs.EnableFuncCallDepth(true)
	logs.Debug("this is a test,my name is %s", "stu01")
	logs.Trace("this is a trace,my name is %s", "stu02")
	logs.Warn("this is a warn,my name is %s", "stu03")


}
