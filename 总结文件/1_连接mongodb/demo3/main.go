package main

import (
	"LearnGo/总结文件/1_连接mongodb/libs/go/data_models/IOCs"
	"LearnGo/总结文件/1_连接mongodb/libs/go/data_models/tags"
	"LearnGo/总结文件/1_连接mongodb/libs/go/data_models/threat"
	"LearnGo/总结文件/1_连接mongodb/libs/go/data_models/threat/alarm"
	"LearnGo/总结文件/1_连接mongodb/libs/go/data_models/threat/attacker"
	"LearnGo/总结文件/1_连接mongodb/libs/go/data_models/threat/evaluate"
	"LearnGo/总结文件/1_连接mongodb/libs/go/data_models/threat/history"
	"LearnGo/总结文件/1_连接mongodb/libs/go/data_models/threat/threat_event"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Student struct{
	Id_ bson.ObjectId `bson:"_id"`
	Name string `bson:"name"`
	Hometown int `bson:"hometown"`
	Age int `bson:"age"`
	Gender bool `bson:"gender"`


}


type MongoAlarm struct {
	Type                          threat.AlarmType                        `json:"type" bson:"type"`
	ThreatID                      string                                  `json:"threat_id" bson:"threat_id"`
	ThreatName                    alarm.ThreatAlarmName                   `json:"name" bson:"name"`
	AlarmSummary                  alarm.AlarmSummary                      `json:"summary" bson:"summary"`
	ConcernedBy                   []string                                `json:"concerned_by" bson:"concerned_by"`
	AlartTag                      []tags.Tags                             `json:"alert_tags" bson:"alert_tags"`
	Evaluations                   evaluate.ThreatEvaluations              `json:"evaluations" bson:"evaluations"`
	Attacker                      attacker.Attacker                       `json:"attacker" bson:"attacker"`
	Related                       alarm.AlarmRelatedInfo                  `json:"related" bson:"related"`
	Asserts                       alarm.RelatedAssert                     `json:"asserts" bson:"asserts"`
	ThreatEvents                  EventList                               `json:"threat_events" bson:"threat_events"`
	HistoricalVulnerabiliyAsserts []history.HistoricalVulnerabiliyAsserts `json:"historical_vulnerabiliy_asserts" bson:"historical_vulnerabiliy_asserts"`
	HistoricalResponseAsserts     []history.HistoricalResponse            `json:"historical_response_asserts" bson:"historical_response_asserts"`
	IoAs                          []IOCs.IIoA                             `json:"IoAs" bson:"IoAs"`
	IoEs                          []IOCs.IIoE                             `json:"IoEs" bson:"IoEs"`
}

type EventList struct {
	Count     int         `json:"count" bson:"count"`
	EventList []mgo.DBRef `json:"eventlist" bson:"eventlist"`
}

func InsertDataToMac(data *MongoAlarm){
	dialInfo := &mgo.DialInfo{
		Addrs: []string{"127.0.0.1:27017"}, //远程(或本地)服务器地址及端口号
		Direct: false,
		Timeout: time.Second * 100,
		Database: "admin", //数据库
		Source: "admin",
		Username: "zhangyanjun",
		Password: "123",
		PoolLimit: 4096, // Session.SetPoolLimit
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c :=session.DB("zyj").C("alarm")
	errs:=c.Insert(data)
	if err!=nil{
		fmt.Println(errs)
	}




}







func main() {
	dialInfo := &mgo.DialInfo{
		Addrs: []string{"10.255.175.225:27017"}, //远程(或本地)服务器地址及端口号
		Direct: false,
		Timeout: time.Second * 100,
		Database: "admin", //数据库
		Source: "admin",
		//Username: "root",
		//Password: "root",
		PoolLimit: 4096, // Session.SetPoolLimit
	}
	session, err := mgo.DialWithInfo(dialInfo)
	fmt.Printf("%T\n",session)
	if err != nil {
		panic(err)
	}
	defer session.Close()



	c :=session.DB("acd").C("alarm")


	//var alarm_list []MongoAlarm

	//cve_list:=[]map[string]string{{"vulnerability_id":"CVE-2014-8361"},{"vulnerability_id":"CVE-2017-17215"}}
	//
    // //err = c.Find(bson.M{"IoEs":bson.M{"vulnerability_id":"CVE-2014-8361"}}).All(&alarm_list )
    // err = c.Find(bson.M{"IoEs":bson.M{"$in":cve_list}}).All(&alarm_list )
	//if err !=nil{
	//	fmt.Println(err)
	//}
	//
	//buf,err :=json.Marshal(alarm_list)
	//fmt.Println(string(buf))



	//查询计数
    //res :=c.Find(bson.M{"type":"threat"})
    //count,err := res.Count()
    //fmt.Print(count)


    //查询一个
	//var alarms MongoAlarm
    //err_one := res.One(&alarms)
    //if err !=nil{
    //	fmt.Println(err_one)
	//}
    //fmt.Println(alarms)

    //查询跳过
	//res :=c.Find(bson.M{"type":"threat"})
    //var alarm_list []MongoAlarm
    //errs :=res.Skip(10).One(&alarm_list )
    //count,_:=res.Skip(10).Count()
    //fmt.Println(count)
    //if errs !=nil{
    //	fmt.Println(errs)
	//}


	//Sort

	//var alarms []MongoAlarm
	//err = c.Find(bson.M{"type":"threat"}).Sort("threat_id").All(&alarms)
	//if err !=nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(alarms)




    db :=session.DB("acd")
    c =db.C("alarm")
    var alarms  MongoAlarm
    err = c.Find(bson.M{"threat_id":"1556615395578259351"}).One(&alarms)
    if err !=nil{
    	fmt.Println(err)
	}

	EventList :=alarms.ThreatEvents.EventList

	refId_list := make([]interface{},0)
	for _, i :=range EventList{
		fmt.Println(i)
		//event := new(threat_event.Event)
		//refEvent:=db.FindRef(&i)
		//err=refEvent.One(event)
		//if err !=nil{
		//	fmt.Println(err)
		//}
		//fmt.Println(event)

		refId := i.Id
		fmt.Println(refId)
		refId_list=append(refId_list,refId)

	}

	fmt.Println(refId_list)

	c_event:=db.C("event")

	var eventlist []threat_event.Event


	res := c_event.Find(bson.M{"_id":bson.M{"$in":refId_list}})

	err =res.Sort("-generated_at").Skip(2).All(&eventlist)

	if err !=nil{
		fmt.Println(err)
	}

	buf,err :=json.Marshal(eventlist)
	fmt.Println(string(buf))

















}
