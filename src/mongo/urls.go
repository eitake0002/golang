/*
 Mongo CRUD package.
*/
package mongo

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const HOST = "mongodb://localhost/panop"
const DB_NAME = "panop"

// Struct for urls collection.
type Url struct {
	Url string
}

// Insert url to urls collection.
func Insert(url string) error {
	session, _ := mgo.Dial(HOST)
	defer session.Close()
	db := session.DB(DB_NAME)

	url_st := Url{Url: url}
	col := db.C("urls")
	err := col.Insert(url_st)
	return err
}

// Read all documents in collection.
func ReadAll() {
	session, _ := mgo.Dial(HOST)
	defer session.Close()
	db := session.DB("urls")

	var results []Url
	err := db.C("urls").Find(bson.M{}).All(&results)
	if err != nil {
		panic(err)
	}
	for _, v := range results {
		log.Println(v)
	}
}
