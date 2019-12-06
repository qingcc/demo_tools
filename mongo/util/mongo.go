package util

import (
	"github.com/globalsign/mgo"
	"sync"
)

var (
	mgoTestConnOnce  sync.Once
	mgoTestDBSession *mgo.Session
	connUrl          = "mongodb://test:123456@localhost:27017/test"
	baseConnUrl      = "mongodb://qing:mongo123@localhost:27017/admin"
)

//[mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]

func GetMgoTestSession() *mgo.Session {
	mgoTestConnOnce.Do(func() {
		url := connUrl
		if session, err := mgo.Dial(url); err != nil {
			panic(err)
		} else {
			mgoTestDBSession = session
			mgoTestDBSession.SetMode(mgo.Monotonic, true)
		}
	})
	return mgoTestDBSession
}

func GetMgoTest(coll string) (c *mgo.Collection) {
	sess := GetMgoTestSession()
	c = sess.DB("test").C(coll)
	return
}
