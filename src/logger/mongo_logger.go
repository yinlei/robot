// 以mongo作为后端存储日志
package logger

import (
	//"fmt"
	"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
)

type MongoLogger struct {
	session *mgo.Session
}

func NewMongoLogger(url string) *MongoLogger {
	mango_session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	return &MongoLogger{session: mango_session}
}
