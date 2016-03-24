package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

var DB *mgo.Database

func init() {
	session, err := mgo.Dial(beego.AppConfig.String("MongoDb"))
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	DB = session.DB("lstank")
}
