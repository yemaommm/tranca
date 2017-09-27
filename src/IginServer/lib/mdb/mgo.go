package mdb

import (
	"gopkg.in/mgo.v2"
	//  "labix.org/v2/mgo/bson"
	"fmt"
	"server/conf"
)

var (
	session      *mgo.Session
	databaseName = conf.GET["config"]["MONGODB_NAME"]
	url          = conf.GET["config"]["MONGODB_URL"]
	username     = conf.GetString("config", "MONGODB_USER")
	password     = conf.GetString("config", "MONGODB_PWD")
)

func Session() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.Dial(url)
		if err != nil {
			panic(err) // no, not really
		}
		session.SetMode(mgo.Eventual, true)
		session.Login(&mgo.Credential{Username: username, Password: password})
	}
	return session.Clone()
}
func M(collection string, f func(*mgo.Collection)) {
	session := Session()
	defer func() {
		session.Close()
		if err := recover(); err != nil {
			fmt.Println("M", err)
		}
	}()
	c := session.DB(databaseName).C(collection)
	f(c)
}
