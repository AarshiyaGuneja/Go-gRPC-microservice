package util

import "gopkg.in/mgo.v2"

func ConnectToMongo() (*mgo.Session, error) {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		return nil, err
	} else {
		return session, nil
	}

}
