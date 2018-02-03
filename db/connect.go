package db

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

var (
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

const (
	// MongoDBUrl is the default mongodb url that will be used to connect to the
	MongoDBUrl = "mongodb://nenjoAdmin:Du1Vyvnwq8f5oeHtDRl3uzApU1FMp24H@localhost:27017/ksusers"
)

// Connect connects to mongodb
func Connect() {
	uri := MongoDBUrl

	if len(uri) == 0 {
		uri = MongoDBUrl
	}

	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)

	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to MongoDB")
	Session = s
	Mongo = mongo
}
