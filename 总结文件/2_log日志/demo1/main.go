package main

import "go.uber.org/zap"

func main() {

	//log := zap.NewExample()
	//fmt.Println("%T",log)
	//log.Debug("this is a Debug message")

	//log,_ := zap.NewDevelopment()
	//log.Debug("this is a Debug message")


	//log,_:= zap.NewProduction()
	//log.Debug("this is a Debug message")
	//log.Info("this is a Info message")


	log,_:=zap.NewDevelopment()
	slogger :=log.Sugar()

	slogger.Debug("hello")
	slogger.Info("hello")

	slogger.Info("Info() uses sprint")
	slogger.Infof("Infof() uses %s", "sprintf")
	slogger.Infow("Infow() allows tags", "name", "Legolas", "type", 1)


}
