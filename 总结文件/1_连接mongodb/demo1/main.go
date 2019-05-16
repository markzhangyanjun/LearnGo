package main

import "gopkg.in/mgo.v2"

func ConnecToDB() *mgo.Collection {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("medex").C("student")
	return c
}
func main() {

	 d := mgo.DialInfo{
		 Addrs :[]string{"127.0.0,1:27017"},
		 Direct : true,
		 Timeout : 5000,
		 Database: "new",



	}

}
