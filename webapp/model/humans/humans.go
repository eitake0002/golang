package humans

import (
  mgo "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "log"
)

type Human struct {
  Name string
  Age  int
}

const host = "mongodb://localhost/humandb"

func Insert() {
  session, _ := mgo.Dial(host)
  defer session.Close()
  db := session.DB("humandb")

  human := Human{"test", 100}
  col := db.C("humans")
  col.Insert(human)
}

func ReadAll() {
  session, _ := mgo.Dial(host)
  defer session.Close()
  db := session.DB("humandb")

  var results []Human
  err := db.C("humans").Find(bson.M{}).All(&results)
  if err != nil {
    panic(err)
  }
  for _, v := range results {
    log.Println(v)
  }
}
